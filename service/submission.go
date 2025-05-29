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
		// UpdateSubmission(ctx context.Context, submission *dto.SubmissionCreateRequest) (*entities.Submission, error)
		DeleteSubmission(ctx context.Context, id uuid.UUID) error
		GetSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]entities.Submission, error)
		GetSubmissionsByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Submission, error)
		GetSubmissionsByAssessmentIDAndUserID(ctx context.Context, assessmentID uuid.UUID, userID uuid.UUID) (*entities.Submission, error)
		GetStudentSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID, flag string)  ([]dto.GetSubmissionStudentResponse, error)
		
	}
	submissionService struct {
		submissionRepo repository.SubmissionRepository
		questionRepo  repository.QuestionRepository
		assessmentRepo repository.AssessmentRepository
	}
)

func NewSubmissionService(submissionRepo repository.SubmissionRepository, questionRepo repository.QuestionRepository, assessmentRepo repository.AssessmentRepository) SubmissionService {
	return &submissionService{
		submissionRepo: submissionRepo,
		questionRepo:  questionRepo,
		assessmentRepo: assessmentRepo,
	}
}

func (submissionService *submissionService) CreateSubmission(ctx context.Context, submission *dto.SubmissionCreateRequest) (dto.SubmissionCreateResponse, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	assesment, errr := submissionService.assessmentRepo.GetAssessmentByID(ctx, nil, submission.AssessmentID)
	if errr != nil {
		return dto.SubmissionCreateResponse{}, errr
	}

	// check if user is a member of the class
	params := url.Values{}
	params.Add("classID", assesment.ClassID.String())
	params.Add("userID", submission.UserID.String())
	url := fmt.Sprintf("%s/service/class/member/?%s", os.Getenv("CLASS_SERVICE_URL"), params.Encode())
	resp, err := http.Get(url)
	if err != nil {
		return dto.SubmissionCreateResponse{}, fmt.Errorf("failed to get class member: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return dto.SubmissionCreateResponse{}, fmt.Errorf("you are not a member of this class")
	}
	defer resp.Body.Close()

	// check if the submission already exists
	_,flag, _ := submissionService.submissionRepo.GetSubmissionsByAssessmentIDAndUserID(ctx, nil, submission.AssessmentID,submission.UserID)
	if flag {
		return dto.SubmissionCreateResponse{}, errors.New("submission already exists")
	}
	submissionEntity := entities.Submission{
		UserID: 	 submission.UserID,
		AssessmentID: submission.AssessmentID,
		Status: "in_progress",
		EndedTime: time.Now().Add(time.Duration(assesment.Duration) * time.Second),
	}
	
	var question []entities.Question
	question, err = submissionService.questionRepo.GetQuestionsByAssessmentID(ctx, nil, submission.AssessmentID)
	if err != nil {
		question = nil
	}

	createdSubmission, err := submissionService.submissionRepo.CreateSubmission(ctx, nil, &submissionEntity)
	if err != nil {
		return dto.SubmissionCreateResponse{}, err
	}
	// create cron job to update submission status
	scheduler := NewSubmissionScheduler(submissionService)
	if err := scheduler.ScheduleAutoSubmit(createdSubmission); err != nil {
		return dto.SubmissionCreateResponse{}, fmt.Errorf("failed to schedule auto submit: %w", err)
	}
	
	// Start the scheduler
	scheduler.Start()

	return dto.SubmissionCreateResponse{
		ID:             createdSubmission.ID,
		UserID:         createdSubmission.UserID,
		AssessmentID:   createdSubmission.AssessmentID,
		EndedTime:      createdSubmission.EndedTime,
		Question: question,
	}, nil
}

func (submissionService *submissionService) GetAllSubmissions(ctx context.Context) ([]entities.Submission, error) {
	if submissions, err := submissionService.submissionRepo.GetAllSubmissions(); err != nil {
		return []entities.Submission{}, err
	} else {
		return submissions, nil
	}
}

func (submissionService *submissionService) GetSubmissionByID(ctx context.Context, id uuid.UUID) (*entities.Submission, error) {
	submission, err := submissionService.submissionRepo.GetSubmissionByID(ctx, nil, id)
	if err != nil {
		return &entities.Submission{}, err
	}
	return submission, nil
}

func (submissionService *submissionService) DeleteSubmission(ctx context.Context, id uuid.UUID) error {
	if err := submissionService.submissionRepo.DeleteSubmission(ctx, nil, id); err != nil {
		return err
	}
	return nil
}

func (submissionService *submissionService) GetSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]entities.Submission, error) {
	submissions, err := submissionService.submissionRepo.GetSubmissionsByAssessmentID(ctx, nil, assessmentID)
	if err != nil {
		return []entities.Submission{}, err
	}
	return submissions, nil
}

func (submissionService *submissionService) GetSubmissionsByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Submission, error) {
	submissions, err := submissionService.submissionRepo.GetSubmissionsByUserID(ctx, nil, userID)
	if err != nil {
		return []entities.Submission{}, err
	}
	return submissions, nil
}

func (submissionService *submissionService) GetSubmissionsByAssessmentIDAndUserID(ctx context.Context, assessmentID uuid.UUID, userID uuid.UUID) (*entities.Submission, error) {
	submission,_, err := submissionService.submissionRepo.GetSubmissionsByAssessmentIDAndUserID(ctx, nil, assessmentID, userID)
	if err != nil {
		return &entities.Submission{}, err
	}
	return submission, nil
}

func (submissionService *submissionService) Submitted(ctx context.Context, submissionID uuid.UUID) (*entities.Submission, error) {
	submission, err := submissionService.submissionRepo.GetSubmissionByID(ctx, nil, submissionID)
	if err != nil {
		return &entities.Submission{}, err
	}
	data := entities.Submission{
		ID: submission.ID,
		UserID: submission.UserID,
		AssessmentID: submission.AssessmentID,
		Status: "submitted",
	}
	result, err := submissionService.submissionRepo.Submitted(ctx, nil, &data)
	if err != nil {
		return &entities.Submission{}, err
	}
	res := entities.Submission{
		ID: result.ID,
		UserID: result.UserID,
		AssessmentID: result.AssessmentID,
		Status: result.Status,
		EndedTime: result.EndedTime,
	}

	return &res, nil
}


func (submissionService *submissionService) GetStudentSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID, flag string) ([]dto.GetSubmissionStudentResponse, error) {
	// Inisialisasi slice dengan kapasitas 0 tapi bukan nil
	studentSubmission := make([]dto.GetSubmissionStudentResponse, 0)
	
	assesment, err := submissionService.assessmentRepo.GetAssessmentByID(ctx, nil, assessmentID)
	if err != nil {
		return studentSubmission, err // Return empty slice instead of nil
	}
	
	submissions, err := submissionService.submissionRepo.GetSubmissionsByAssessmentID(ctx, nil, assessmentID)
	if err != nil {
		return studentSubmission, err // Return empty slice instead of nil
	}
	
	classID := assesment.ClassID
	if err := godotenv.Load(); err != nil {
		return studentSubmission, fmt.Errorf("failed to load environment variables: %w", err)
	}

	url := fmt.Sprintf("%s/service/class/%s", os.Getenv("CLASS_SERVICE_URL"), classID)
	resp, err := http.Get(url)
	if err != nil {
		return studentSubmission, err
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return studentSubmission, fmt.Errorf("failed to read response body: %w", err)
	}
	
	var members []dto.GetSubmissionStudentResponse
	if err := json.Unmarshal(body, &members); err != nil {
		return studentSubmission, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	
	// Buat map untuk submission
	submissionMap := make(map[uuid.UUID]entities.Submission)
	for _, s := range submissions {
		submissionMap[s.UserID] = s
	}
	
	defaultStatus := entities.ExamStatus("todo")
	
	for _, m := range members {
		// Skip teacher
		if m.Role == dto.RoleTeacher {
			continue
		}
		
		// Cek apakah ada submission untuk user ini
		sub, hasSubmission := submissionMap[m.User_userID]
		
		var datas dto.GetSubmissionStudentResponse
		datas.Username = m.Username
		datas.User_userID = m.User_userID
		datas.Kelas_kelasID = m.Kelas_kelasID
		datas.Role = m.Role
		
		if !hasSubmission {
			// Tidak ada submission, set default values
			datas.ID = nil
			datas.Status = defaultStatus
			datas.TimeRemaining = nil
			datas.Score = 0
		} else {
			// Ada submission, set berdasarkan status
			switch sub.Status {
			case "in_progress":
				now := time.Now()
				duration := int64(sub.EndedTime.Sub(now).Seconds())
				if duration < 0 {
					duration = 0
				}
				
				datas.ID = &sub.ID
				datas.Status = sub.Status
				datas.TimeRemaining = &duration
				datas.Score = sub.Score
				
			case "submitted":
				datas.ID = &sub.ID
				datas.Status = sub.Status
				datas.TimeRemaining = nil
				datas.Score = sub.Score
				
			case "todo":
				datas.ID = nil
				datas.Status = sub.Status
				datas.TimeRemaining = nil
				datas.Score = 0
				
			default:
				// Status tidak dikenal, gunakan default
				datas.ID = nil
				datas.Status = defaultStatus
				datas.TimeRemaining = nil
				datas.Score = 0
			}
		}
		
		// Filter berdasarkan flag
		if flag == "" || string(datas.Status) == flag {
			studentSubmission = append(studentSubmission, datas)
		}
	}
	
	return studentSubmission, nil
}



func (s *SubmissionScheduler) Start() {
	s.cron.Start()
}

func (s *SubmissionScheduler) Stop() {
	s.cron.Stop()
}

func NewSubmissionScheduler(submissionService *submissionService) *SubmissionScheduler {
	return &SubmissionScheduler{
		cron: cron.New(),
		submissionService: submissionService,
		jobs: make(map[string]cron.EntryID),
	}
}

// Dalam CreateSubmission method
type SubmissionScheduler struct {
	cron *cron.Cron
	submissionService *submissionService
	jobs map[string]cron.EntryID // Map untuk track job IDs
	mu   sync.RWMutex           // Mutex untuk thread safety
}

func (s *SubmissionScheduler) ScheduleAutoSubmit(submission *entities.Submission) error {
	// Cron expression format: second minute hour day month dayofweek
	// Tapi robfig/cron/v3 menggunakan format standar: minute hour day month dayofweek
	cronExpr := fmt.Sprintf("%d %d %d %d *", 
		submission.EndedTime.Minute(),
		submission.EndedTime.Hour(),
		submission.EndedTime.Day(),
		int(submission.EndedTime.Month()),
	)

	jobKey := fmt.Sprintf("auto_submit_%s", submission.ID.String())
	
	entryID, err := s.cron.AddFunc(cronExpr, func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
		defer cancel()
		
		s.submissionService.Submitted(ctx, submission.ID)
		
		// Remove job setelah dijalankan
		s.removeJob(jobKey)
	})
	
	if err != nil {
		return err
	}
	
	// Simpan job ID untuk bisa di-remove nanti
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

// Method untuk cancel job jika submission di-submit manual sebelum deadline
func (s *SubmissionScheduler) CancelAutoSubmit(submissionID string) {
	jobKey := fmt.Sprintf("auto_submit_%s", submissionID)
	s.removeJob(jobKey)
}

// Alternatif: Background service yang cek database secara berkala
func (submissionService *submissionService) StartAutoSubmitWorker() {
	ticker := time.NewTicker(time.Minute * 1) // Cek setiap menit
	go func() {
		for range ticker.C {
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
			
			// Get submissions yang sudah expired tapi belum submitted
			expiredSubmissions, err := submissionService.submissionRepo.GetExpiredSubmissions(ctx, nil)
			if err != nil {
				fmt.Printf("Failed to get expired submissions: %v\n", err)
				cancel()
				continue
			}
			
			for _, submission := range expiredSubmissions {
				submissionService.Submitted(ctx, submission.ID)
			}
			
			cancel()
		}
	}()
}