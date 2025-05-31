package service

import (
	dto "assesment/dto"
	entities "assesment/entities"
	repository "assesment/repository"
	"assesment/utils"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	AnswerService interface {
		CreateAnswer(ctx context.Context, answer *dto.AnswerCreateRequest) (dto.AnswerResponse, error)
		GetAllAnswers(ctx context.Context) ([]entities.Answer, error)
		GetAnswerByID(ctx context.Context, id uuid.UUID) (*entities.Answer, error)
		UpdateAnswer(ctx context.Context, answer *dto.AnswerUpdateRequest) (*dto.AnswerUpdatedResponse, error)
		GetAnswerByQuestionID(ctx context.Context, questionID uuid.UUID) ([]entities.Answer, error)
		GetAnswerBySubmissionID(ctx context.Context, submissionID uuid.UUID) ([]dto.GetAnswerBySubmissionIDResponse, error)
		// GetAnswerByStudentID(ctx context.Context, id dto.GetAnswerByStudentIDRequest) ([]entities.Answer, error)
	}
	answerService struct {
		answerRepo     repository.AnswerRepository
		submissionRepo repository.SubmissionRepository
		assesmentRepo repository.AssessmentRepository
	}
)

func NewAnswerService(answerRepo repository.AnswerRepository,submissionRepo repository.SubmissionRepository, assesmentRepo repository.AssessmentRepository) AnswerService {
	return &answerService{
		answerRepo: answerRepo,
		submissionRepo: submissionRepo,
		assesmentRepo: assesmentRepo,
	}
}

func (answerService *answerService) CreateAnswer(ctx context.Context, answer *dto.AnswerCreateRequest) (dto.AnswerResponse, error) {
	// Check if submission exists and is in progress
	submission, err := answerService.submissionRepo.GetSubmissionByID(ctx, nil, answer.SubmisiionID)
	if err != nil {
		return dto.AnswerResponse{}, utils.ErrCreateAnswer
	}
	if submission.Status != "in_progress" {
		return dto.AnswerResponse{}, errors.New("submission is not in progress")
	}
	if submission.EndedTime.Before(time.Now()) {
		return dto.AnswerResponse{}, errors.New("submission has ended")
	}
	
	assement, err := answerService.assesmentRepo.GetAssessmentByID(ctx, nil, submission.AssessmentID)
	if err != nil {
		return dto.AnswerResponse{}, utils.ErrCreateAnswer
	}
	if time.Now().After(assement.EndTime) {
		return dto.AnswerResponse{}, errors.New("assesment has ended")
	}
	
	answerExists, err := answerService.answerRepo.GetAnswerBySubmissionIDAndQuestionID(ctx, nil, answer.SubmisiionID, answer.IdQuestion)
	if err != nil {
		return dto.AnswerResponse{}, utils.ErrCreateAnswer
	}
	if answerExists.ID != uuid.Nil {
		return dto.AnswerResponse{}, errors.New("answer already exists")
	}
	
	answerEntity := entities.Answer{
		QuestionID: answer.IdQuestion,
		SubmissionID: answer.SubmisiionID,
		ChoiceID: answer.IdChoice,
	}
	
	createdAnswer, err := answerService.answerRepo.CreateAnswer(ctx, nil, &answerEntity)
	if err != nil {
		return dto.AnswerResponse{}, utils.ErrCreateAnswer
	}
	return createdAnswer, nil
}

func (answerService *answerService) GetAllAnswers(ctx context.Context) ([]entities.Answer, error) {
	if answers, err := answerService.answerRepo.GetAllAnswers(); err != nil {
		return []entities.Answer{}, utils.ErrGetAllAnswers
	} else {
		return answers, nil
	}
}

func (answerService *answerService) GetAnswerByID(ctx context.Context, id uuid.UUID) (*entities.Answer, error) {
	answer, err := answerService.answerRepo.GetAnswerByID(ctx, nil, id)
	if err != nil {
		return &entities.Answer{}, err
	}
	return answer, nil	
}

func (answerService *answerService) UpdateAnswer(ctx context.Context, answer *dto.AnswerUpdateRequest) (*dto.AnswerUpdatedResponse, error) {
	answ,err := answerService.answerRepo.GetAnswerByID(ctx, nil, answer.ID)
	if err != nil {
		return nil, utils.ErrGetAnswerByID
	}
	if answ == nil {
		return &dto.AnswerUpdatedResponse{}, utils.ErrGetAnswerByID
	}
	if answ.Submission.Status != "in_progress" {
		return &dto.AnswerUpdatedResponse{}, errors.New("submission is not in progress")
	}
	data := entities.Answer{
		ID: answ.ID,
		QuestionID: answer.IdQuestion,
		SubmissionID: answ.SubmissionID,
		ChoiceID: answer.IdChoice,
	}
	updatedAnswer, err := answerService.answerRepo.UpdateAnswer(ctx, nil, &data)
	if err != nil {
		return &dto.AnswerUpdatedResponse{}, utils.ErrUpdateAnswer
	}
	return &dto.AnswerUpdatedResponse{
		ID: updatedAnswer.ID,
		QuestionID: updatedAnswer.QuestionID,
		ChoiceID: updatedAnswer.ChoiceID,
		SubmissionID: updatedAnswer.SubmissionID,
		CreatedAt: updatedAnswer.CreatedAt,
		UpdatedAt: updatedAnswer.UpdatedAt,
		DeletedAt: updatedAnswer.DeletedAt,
	}, nil
}

func (answerService *answerService) GetAnswerByQuestionID(ctx context.Context, questionID uuid.UUID) ([]entities.Answer, error) {
	answers, err := answerService.answerRepo.GetAnswerByQuestionID(ctx, nil, questionID)
	if err != nil {
		return []entities.Answer{}, utils.ErrGetAnswerByQuestionID
	}
	return answers, nil
}

func (answerService *answerService) GetAnswerBySubmissionID(ctx context.Context, submissionID uuid.UUID) ([]dto.GetAnswerBySubmissionIDResponse, error) {
	answer, err := answerService.answerRepo.GetAnswerBySubmissionID(ctx, nil, submissionID)
	if err != nil {
		return []dto.GetAnswerBySubmissionIDResponse{}, utils.ErrGetAnswerBySubmissionID
	}
	return answer, nil
}	