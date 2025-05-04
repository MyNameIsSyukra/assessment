package service

import (
	dto "assesment/dto"
	entities "assesment/entities"
	repository "assesment/repository"
	"context"
	"errors"
	"fmt"

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
		GetSubmissionsByAssessmentIDAndClassID(ctx context.Context, assessmentID uuid.UUID, classID uuid.UUID) ([]entities.Submission, error)
		Submitted(ctx context.Context, submissionID uuid.UUID) (*entities.Submission, error)
	}
	submissionService struct {
		submissionRepo repository.SubmissionRepository
		questionRepo  repository.QuestionRepository
	}
)

func NewSubmissionService(submissionRepo repository.SubmissionRepository, questionRepo repository.QuestionRepository) SubmissionService {
	return &submissionService{
		submissionRepo: submissionRepo,
		questionRepo:  questionRepo,
	}
}

func (submissionService *submissionService) CreateSubmission(ctx context.Context, submission *dto.SubmissionCreateRequest) (dto.SubmissionCreateResponse, error) {
	var question []entities.Question
	_,flag, _ := submissionService.submissionRepo.GetSubmissionsByAssessmentIDAndUserID(ctx, nil, submission.AssessmentID,submission.UserID)
	// check if the submission already exists
	if flag {
		return dto.SubmissionCreateResponse{}, errors.New("submission already exists")
	}
	submissionEntity := entities.Submission{
		UserID: 	 submission.UserID,
		AssessmentID: submission.AssessmentID,
		Status: "in_progress",
	}

	fmt.Print(submissionEntity)

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

func (submissionService *submissionService) GetSubmissionsByAssessmentIDAndClassID(ctx context.Context, assessmentID uuid.UUID, classID uuid.UUID) ([]entities.Submission, error) {
	submissions, err := submissionService.submissionRepo.GetSubmissionsByAssessmentIDAndClassID(ctx, nil, assessmentID, classID)
	if err != nil {
		return nil, err
	}
	return submissions, nil
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


