package service

import (
	dto "assesment/dto"
	entities "assesment/entities"
	repository "assesment/repository"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

type (
	SubmissionService interface {
		// student
		CreateSubmission(ctx context.Context, submission *dto.SubmissionCreateRequest) (dto.SubmissionCreateResponse, error)
		Submitted(ctx context.Context, submissionID uuid.UUID) (*entities.Submission, error)
		GetAllSubmissions(ctx context.Context) ([]entities.Submission, error)
		GetSubmissionByID(ctx context.Context, id uuid.UUID) (*entities.Submission, error)
		DeleteSubmission(ctx context.Context, id uuid.UUID) error
		GetSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]entities.Submission, error)
		GetSubmissionsByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Submission, error)
		GetSubmissionsByAssessmentIDAndUserID(ctx context.Context, assessmentID uuid.UUID, userID uuid.UUID) (*entities.Submission, error)
		GetStudentSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID, flag string) ([]dto.GetSubmissionStudentResponse, error)
		
		// Add cleanup method
		Close() error
	}
	
	submissionService struct {
		submissionRepo repository.SubmissionRepository
		questionRepo   repository.QuestionRepository
		assessmentRepo repository.AssessmentRepository
		scheduler      *SubmissionScheduler
		httpClient     *http.Client
		classServiceURL string
		workerCtx      context.Context
		workerCancel   context.CancelFunc
	}
)

// Global scheduler instance - singleton pattern
var (
	globalScheduler *SubmissionScheduler
	schedulerOnce   sync.Once
)

func getGlobalScheduler(service *submissionService) *SubmissionScheduler {
	schedulerOnce.Do(func() {
		globalScheduler = &SubmissionScheduler{
			cron:              cron.New(cron.WithSeconds()), // Support seconds for precise timing
			submissionService: service,
			jobs:              make(map[string]cron.EntryID),
		}
		globalScheduler.cron.Start()
	})
	return globalScheduler
}

func NewSubmissionService(submissionRepo repository.SubmissionRepository, questionRepo repository.QuestionRepository, assessmentRepo repository.AssessmentRepository) SubmissionService {
	// Load environment variables once
	if err := godotenv.Load(); err != nil {
		// Log error but don't panic in production
		fmt.Printf("Warning: failed to load .env file: %v\n", err)
	}

	service := &submissionService{
		submissionRepo:  submissionRepo,
		questionRepo:    questionRepo,
		assessmentRepo:  assessmentRepo,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
		classServiceURL: os.Getenv("CLASS_SERVICE_URL"),
	}
	
	service.scheduler = getGlobalScheduler(service)
	
	// Initialize context for background worker
	service.workerCtx, service.workerCancel = context.WithCancel(context.Background())
	
	// Start background worker as backup mechanism
	service.StartAutoSubmitWorker(service.workerCtx)
	
	return service
}

func (s *submissionService) Close() error {
	// Cancel background worker
	if s.workerCancel != nil {
		s.workerCancel()
	}
	
	// Stop scheduler
	if s.scheduler != nil {
		s.scheduler.Stop()
	}
	return nil
}

func (s *submissionService) CreateSubmission(ctx context.Context, submission *dto.SubmissionCreateRequest) (dto.SubmissionCreateResponse, error) {
	// Get assessment details
	assessment, err := s.assessmentRepo.GetAssessmentByID(ctx, nil, submission.AssessmentID)
	if err != nil {
		return dto.SubmissionCreateResponse{}, fmt.Errorf("failed to get assessment: %w", err)
	}

	// Check if user is a member of the class
	if err := s.checkClassMembership(ctx, assessment.ClassID, submission.UserID); err != nil {
		return dto.SubmissionCreateResponse{}, err
	}
	// Check if assessment is active
	if time.Now().Before(assessment.StartTime) || time.Now().After(assessment.EndTime) {
		return dto.SubmissionCreateResponse{}, errors.New("assessment is not active")
	}
	// Check if submission already exists
	_, exists, _ := s.submissionRepo.GetSubmissionsByAssessmentIDAndUserID(ctx, nil, submission.AssessmentID, submission.UserID)
	if exists {
		return dto.SubmissionCreateResponse{}, errors.New("submission already exists")
	}

	// Create submission entity
	endTime := time.Now().Add(time.Duration(assessment.Duration) * time.Second)
	submissionEntity := entities.Submission{
		UserID:       submission.UserID,
		AssessmentID: submission.AssessmentID,
		Status:       "in_progress",
		EndedTime:    endTime,
	}

	// Get questions
	questions, err := s.questionRepo.GetQuestionsByAssessmentID(ctx, nil, submission.AssessmentID)
	if err != nil {
		// Log error but don't fail the submission creation
		fmt.Printf("Warning: failed to get questions: %v\n", err)
		questions = nil
	}


	// Create submission
	createdSubmission, err := s.submissionRepo.CreateSubmission(ctx, nil, &submissionEntity)
	if err != nil {
		return dto.SubmissionCreateResponse{}, fmt.Errorf("failed to create submission: %w", err)
	}

	// Schedule auto-submit - use the global scheduler
	if err := s.scheduler.ScheduleAutoSubmit(createdSubmission); err != nil {
		// Log error but don't fail the submission creation
		fmt.Printf("Warning: failed to schedule auto submit: %v\n", err)
	}

	return dto.SubmissionCreateResponse{
		ID:           createdSubmission.ID,
		UserID:       createdSubmission.UserID,
		AssessmentID: createdSubmission.AssessmentID,
		EndedTime:    createdSubmission.EndedTime,
		Question:     dto.ToQuestionResponses(questions),
	}, nil
}

func (s *submissionService) checkClassMembership(ctx context.Context, classID, userID uuid.UUID) error {
	params := url.Values{}
	params.Add("classID", classID.String())
	params.Add("userID", userID.String())
	
	reqURL := fmt.Sprintf("%s/service/class/member/?%s", s.classServiceURL, params.Encode())
	
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to check class membership: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("you are not a member of this class")
	}
	
	return nil
}

func (s *submissionService) GetAllSubmissions(ctx context.Context) ([]entities.Submission, error) {
	submissions, err := s.submissionRepo.GetAllSubmissions()
	if err != nil {
		return nil, fmt.Errorf("failed to get all submissions: %w", err)
	}
	return submissions, nil
}

func (s *submissionService) GetSubmissionByID(ctx context.Context, id uuid.UUID) (*entities.Submission, error) {
	submission, err := s.submissionRepo.GetSubmissionByID(ctx, nil, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get submission by ID: %w", err)
	}
	return submission, nil
}

func (s *submissionService) DeleteSubmission(ctx context.Context, id uuid.UUID) error {
	// Cancel any scheduled auto-submit for this submission
	s.scheduler.CancelAutoSubmit(id.String())
	
	if err := s.submissionRepo.DeleteSubmission(ctx, nil, id); err != nil {
		return fmt.Errorf("failed to delete submission: %w", err)
	}
	return nil
}

func (s *submissionService) GetSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]entities.Submission, error) {
	submissions, err := s.submissionRepo.GetSubmissionsByAssessmentID(ctx, nil, assessmentID)
	if err != nil {
		return nil, fmt.Errorf("failed to get submissions by assessment ID: %w", err)
	}
	return submissions, nil
}

func (s *submissionService) GetSubmissionsByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Submission, error) {
	submissions, err := s.submissionRepo.GetSubmissionsByUserID(ctx, nil, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get submissions by user ID: %w", err)
	}
	return submissions, nil
}

func (s *submissionService) GetSubmissionsByAssessmentIDAndUserID(ctx context.Context, assessmentID uuid.UUID, userID uuid.UUID) (*entities.Submission, error) {
	submission, _, err := s.submissionRepo.GetSubmissionsByAssessmentIDAndUserID(ctx, nil, assessmentID, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get submission by assessment and user ID: %w", err)
	}
	return submission, nil
}

func (s *submissionService) Submitted(ctx context.Context, submissionID uuid.UUID) (*entities.Submission, error) {
	submission, err := s.submissionRepo.GetSubmissionByID(ctx, nil, submissionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get submission: %w", err)
	}

	// Don't update if already submitted
	if submission.Status == "submitted" {
		return submission, nil
	}

	data := entities.Submission{
		ID:           submission.ID,
		UserID:       submission.UserID,
		AssessmentID: submission.AssessmentID,
		SubmittedAt: time.Now(),
		Status:       "submitted",
	}
	
	result, err := s.submissionRepo.Submitted(ctx, nil, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to update submission status: %w", err)
	}

	// Cancel auto-submit job
	s.scheduler.CancelAutoSubmit(result.ID.String())

	return result, nil
}

func (s *submissionService) GetStudentSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID, flag string) ([]dto.GetSubmissionStudentResponse, error) {
	studentSubmission := make([]dto.GetSubmissionStudentResponse, 0)
	
	assessment, err := s.assessmentRepo.GetAssessmentByID(ctx, nil, assessmentID)
	if err != nil {
		return studentSubmission, fmt.Errorf("failed to get assessment: %w", err)
	}
	
	submissions, err := s.submissionRepo.GetSubmissionsByAssessmentID(ctx, nil, assessmentID)
	if err != nil {
		return studentSubmission, fmt.Errorf("failed to get submissions: %w", err)
	}
	
	// Get class members
	members, err := s.getClassMembers(ctx, assessment.ClassID)
	if err != nil {
		return studentSubmission, fmt.Errorf("failed to get class members: %w", err)
	}
	
	// Create submission lookup map
	submissionMap := make(map[uuid.UUID]entities.Submission)
	for _, s := range submissions {
		submissionMap[s.UserID] = s
	}
	
	defaultStatus := entities.ExamStatus("todo")
	
	for _, member := range members {
		// Skip teachers
		if member.Role == dto.RoleTeacher {
			continue
		}
		
		response := dto.GetSubmissionStudentResponse{
			Username:      member.Username,
			User_userID:   member.User_userID,
			Kelas_kelasID: member.Kelas_kelasID,
			Role:          member.Role,
			PhotoUrl:      fmt.Sprintf("%s/storage/user_profile_pictures/%s.jpg",os.Getenv("GATEWAY_URL"), member.User_userID.String()),
		}
		
		if submission, hasSubmission := submissionMap[member.User_userID]; hasSubmission {
			s.populateSubmissionResponse(&response, &submission)
		} else {
			// No submission found
			response.ID = nil
			response.Status = defaultStatus
			response.TimeRemaining = &assessment.Duration
			response.Score = 0
		}
		
		// Apply filter
		if flag == "" || string(response.Status) == flag {
			studentSubmission = append(studentSubmission, response)
		}
	}
	
	return studentSubmission, nil
}

func (s *submissionService) getClassMembers(ctx context.Context, classID uuid.UUID) ([]dto.GetSubmissionStudentResponse, error) {
	reqURL := fmt.Sprintf("%s/service/class/%s", s.classServiceURL, classID)
	
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get class members: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get class members: status %d", resp.StatusCode)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	
	var members []dto.GetSubmissionStudentResponse
	if err := json.Unmarshal(body, &members); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	
	return members, nil
}

func (s *submissionService) populateSubmissionResponse(response *dto.GetSubmissionStudentResponse, submission *entities.Submission) {
	switch submission.Status {
	case "in_progress":
		now := time.Now()
		duration := int(submission.EndedTime.Sub(now).Seconds())
		if duration < 0 {
			duration = 0
		}
		
		response.ID = &submission.ID
		response.Status = submission.Status
		response.TimeRemaining = &duration
		response.Score = int(submission.Score)
		
	case "submitted":
		response.ID = &submission.ID
		response.Status = submission.Status
		response.TimeRemaining = nil
		response.Score = int(submission.Score)
		
	case "todo":
		response.ID = nil
		response.Status = submission.Status
		response.TimeRemaining = nil
		response.Score = 0
		
	default:
		// Unknown status, use defaults
		response.ID = nil
		response.Status = entities.ExamStatus("todo")
		response.TimeRemaining = nil
		response.Score = 0
	}
}

// SubmissionScheduler with improved implementation
type SubmissionScheduler struct {
	cron              *cron.Cron
	submissionService *submissionService
	jobs              map[string]cron.EntryID
	mu                sync.RWMutex
}

func (s *SubmissionScheduler) ScheduleAutoSubmit(submission *entities.Submission) error {
	// Calculate time until submission ends
	now := time.Now()
	if submission.EndedTime.Before(now) {
		// Already expired, submit immediately
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
			defer cancel()
			s.submissionService.Submitted(ctx, submission.ID)
		}()
		return nil
	}
	
	// Use AddFunc with @at syntax for one-time execution
	jobKey := fmt.Sprintf("auto_submit_%s", submission.ID.String())
	
	// Calculate seconds until deadline
	secondsUntil := int(submission.EndedTime.Sub(now).Seconds())
	
	entryID, err := s.cron.AddFunc(fmt.Sprintf("@every %ds", secondsUntil), func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
		defer cancel()
		
		// Auto-submit the submission
		if _, err := s.submissionService.Submitted(ctx, submission.ID); err != nil {
			fmt.Printf("Failed to auto-submit submission %s: %v\n", submission.ID, err)
		}
		
		// Remove the job after execution
		s.removeJob(jobKey)
	})
	
	if err != nil {
		return fmt.Errorf("failed to schedule auto submit: %w", err)
	}
	
	// Store job ID
	s.mu.Lock()
	s.jobs[jobKey] = entryID
	s.mu.Unlock()
	
	return nil
}

func (s *SubmissionScheduler) removeJob(jobKey string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if entryID, exists := s.jobs[jobKey]; exists {
		s.cron.Remove(entryID)
		delete(s.jobs, jobKey)
	}
}

func (s *SubmissionScheduler) CancelAutoSubmit(submissionID string) {
	jobKey := fmt.Sprintf("auto_submit_%s", submissionID)
	s.removeJob(jobKey)
}

func (s *SubmissionScheduler) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	// Remove all jobs
	for jobKey, entryID := range s.jobs {
		s.cron.Remove(entryID)
		delete(s.jobs, jobKey)
	}
	
	s.cron.Stop()
}

// Background worker as alternative/backup
func (s *submissionService) StartAutoSubmitWorker(ctx context.Context) {
	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		defer ticker.Stop()
		
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.processExpiredSubmissions()
			}
		}
	}()
}

func (s *submissionService) processExpiredSubmissions() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()
	
	expiredSubmissions, err := s.submissionRepo.GetExpiredSubmissions(ctx, nil)
	if err != nil {
		fmt.Printf("Failed to get expired submissions: %v\n", err)
		return
	}
	
	for _, submission := range expiredSubmissions {
		if _, err := s.Submitted(ctx, submission.ID); err != nil {
			fmt.Printf("Failed to auto-submit expired submission %s: %v\n", submission.ID, err)
		}
	}
}