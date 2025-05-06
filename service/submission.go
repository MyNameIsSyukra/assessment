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
	"time"

	"github.com/google/uuid"
)

type (
	SubmissionService interface {
		CreateSubmission(ctx context.Context, submission *dto.SubmissionCreateRequest) (dto.SubmissionCreateResponse, error)
		GetAllSubmissions(ctx context.Context) ([]entities.Submission, error)
		GetSubmissionByID(ctx context.Context, id uuid.UUID) (*entities.Submission, error)
		// UpdateSubmission(ctx context.Context, submission *dto.SubmissionCreateRequest) (*entities.Submission, error)
		DeleteSubmission(ctx context.Context, id uuid.UUID) error
		GetSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]entities.Submission, error)
		GetSubmissionsByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Submission, error)
		GetSubmissionsByAssessmentIDAndUserID(ctx context.Context, assessmentID uuid.UUID, userID uuid.UUID) (*entities.Submission, error)
		Submitted(ctx context.Context, submissionID uuid.UUID) (*entities.Submission, error)
		GetStudentSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID, flag string) ([]dto.GetSubmissionStudentResponse, error)
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
	var question []entities.Question
	assesment, errr := submissionService.assessmentRepo.GetAssessmentByID(ctx, nil, submission.AssessmentID)
	if errr != nil {
		return dto.SubmissionCreateResponse{}, errr
	}
	_,flag, _ := submissionService.submissionRepo.GetSubmissionsByAssessmentIDAndUserID(ctx, nil, submission.AssessmentID,submission.UserID)
	// check if the submission already exists
	if flag {
		return dto.SubmissionCreateResponse{}, errors.New("submission already exists")
	}
	submissionEntity := entities.Submission{
		UserID: 	 submission.UserID,
		AssessmentID: submission.AssessmentID,
		Status: "in_progress",
		EndedTime: time.Now().Add(time.Duration(assesment.Duration) * time.Second),
	}

	question, err := submissionService.questionRepo.GetQuestionsByAssessmentID(ctx, nil, submission.AssessmentID)
	if err != nil {
		question = nil
	}

	createdSubmission, err := submissionService.submissionRepo.CreateSubmission(ctx, nil, &submissionEntity)
	if err != nil {
		return dto.SubmissionCreateResponse{}, err
	}

	return dto.SubmissionCreateResponse{
		ID:             createdSubmission.ID,
		UserID:         createdSubmission.UserID,
		AssessmentID:   createdSubmission.AssessmentID,
		Question: question,
	}, nil
}

func (submissionService *submissionService) GetAllSubmissions(ctx context.Context) ([]entities.Submission, error) {
	if submissions, err := submissionService.submissionRepo.GetAllSubmissions(); err != nil {
		return nil, err
	} else {
		return submissions, nil
	}
}

func (submissionService *submissionService) GetSubmissionByID(ctx context.Context, id uuid.UUID) (*entities.Submission, error) {
	submission, err := submissionService.submissionRepo.GetSubmissionByID(ctx, nil, id)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return submissions, nil
}

func (submissionService *submissionService) GetSubmissionsByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Submission, error) {
	submissions, err := submissionService.submissionRepo.GetSubmissionsByUserID(ctx, nil, userID)
	if err != nil {
		return nil, err
	}
	return submissions, nil
}

func (submissionService *submissionService) GetSubmissionsByAssessmentIDAndUserID(ctx context.Context, assessmentID uuid.UUID, userID uuid.UUID) (*entities.Submission, error) {
	submission,_, err := submissionService.submissionRepo.GetSubmissionsByAssessmentIDAndUserID(ctx, nil, assessmentID, userID)
	if err != nil {
		return nil, err
	}
	return submission, nil
}

func (submissionService *submissionService) Submitted(ctx context.Context, submissionID uuid.UUID) (*entities.Submission, error) {
	submission, err := submissionService.submissionRepo.GetSubmissionByID(ctx, nil, submissionID)
	if err != nil {
		return nil, err
	}
	data := entities.Submission{
		ID: submission.ID,
		UserID: submission.UserID,
		AssessmentID: submission.AssessmentID,
		Status: "submitted",
	}
	result, err := submissionService.submissionRepo.Submitted(ctx, nil, &data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (submissionService *submissionService) GetStudentSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID,flag string) ([]dto.GetSubmissionStudentResponse, error) {
	class, err := submissionService.assessmentRepo.GetAssessmentByID(ctx, nil, assessmentID)
	if err != nil {
		return nil, err
	}
	submissions, err := submissionService.submissionRepo.GetSubmissionsByAssessmentID(ctx, nil, assessmentID)
	if err != nil {
		return nil, err
	}
	classID := class.ClassID
	url := fmt.Sprintf("http://localhost:8081/service/class/%s", classID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	var members []dto.GetSubmissionStudentResponse
	if err := json.Unmarshal(body, &members); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	
	submissionMap := make(map[uuid.UUID]entities.Submission)
	for _, s := range submissions {
		// Gunakan user_id sebagai key (kalau ada banyak submission per user, bisa dihandle pakai slice)
		submissionMap[s.UserID] = s
	}
	
	status := entities.ExamStatus("todo")
	var result []dto.GetSubmissionStudentResponse
	for _, m := range members {
		if m.Role == dto.RoleTeacher{
			continue
		}
		sub := submissionMap[m.User_userID]
		if sub.Status == "" {
			sub.Status = status
			m.ID = uuid.Nil
		}
		dto := dto.GetSubmissionStudentResponse{
			ID:           m.ID,
			Username:     m.Username,
			User_userID:   m.User_userID,
			Kelas_kelasID: m.Kelas_kelasID,
			Role:         m.Role,
			Status:       sub.Status,
			Score:        sub.Score,
			CreatedAt:    m.CreatedAt,
			UpdatedAt:    m.UpdatedAt,
			DeletedAt:    m.DeletedAt,
		}
		result = append(result, dto)
	}
	var res []dto.GetSubmissionStudentResponse
	if flag == "" {
		return result, nil
	}
	
	for _, m := range result {
		if m.Status == entities.ExamStatus(flag){
			res = append(res, m)
		}
	}
	return res, nil
}


