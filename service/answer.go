package service

import (
	dto "assesment/dto"
	entities "assesment/entities"
	repository "assesment/repository"
	"context"
)

type (
	AnswerService interface {
		CreateAnswer(ctx context.Context, answer *dto.AnswerCreateRequest) (dto.AnswerResponse, error)
		GetAllAnswers(ctx context.Context) ([]entities.Answer, error)
		GetAnswerByID(ctx context.Context, id string) (*entities.Answer, error)
		UpdateAnswer(ctx context.Context, answer *dto.AnswerUpdateRequest) (*entities.Answer, error)
		GetAnswerByQuestionID(ctx context.Context, questionID string) ([]entities.Answer, error)
		GetAnswerByStudentID(ctx context.Context, id dto.GetAnswerByStudentIDRequest) ([]entities.Answer, error)
	}
	answerService struct {
		answerRepo repository.AnswerRepository
	}
)

func NewAnswerService(answerRepo repository.AnswerRepository) AnswerService {
	return &answerService{
		answerRepo: answerRepo,
	}
}

func (answerService *answerService) CreateAnswer(ctx context.Context, answer *dto.AnswerCreateRequest) (dto.AnswerResponse, error) {
	answerEntity := entities.Answer{
		QuestionID: answer.IdQuestion,
		SubmissionID: answer.SubmisiionID,
		ChoiceID: answer.IdChoice,
	}

	createdAnswer, err := answerService.answerRepo.CreateAnswer(ctx, nil, &answerEntity)
	if err != nil {
		return dto.AnswerResponse{}, err
	}
	return createdAnswer, nil
}

func (answerService *answerService) GetAllAnswers(ctx context.Context) ([]entities.Answer, error) {
	if answers, err := answerService.answerRepo.GetAllAnswers(); err != nil {
		return nil, err
	} else {
		return answers, nil
	}
}

func (answerService *answerService) GetAnswerByID(ctx context.Context, id string) (*entities.Answer, error) {
	answer, err := answerService.answerRepo.GetAnswerByID(ctx, nil, id)
	if err != nil {
		return nil, err
	}
	return answer, nil	
}

func (answerService *answerService) UpdateAnswer(ctx context.Context, answer *dto.AnswerUpdateRequest) (*entities.Answer, error) {
	answ,err := answerService.answerRepo.GetAnswerByID(ctx, nil, answer.ID)
	if answ == nil {
		return nil, err
	}
	data := entities.Answer{
		ID: answ.ID,
		QuestionID: answer.IdQuestion,
		SubmissionID: answ.SubmissionID,
		ChoiceID: answer.IdChoice,
	}

	updatedAnswer, err := answerService.answerRepo.UpdateAnswer(ctx, nil, &data)
	if err != nil {
		return nil, err
	}
	return updatedAnswer, nil
}

func (answerService *answerService) GetAnswerByQuestionID(ctx context.Context, questionID string) ([]entities.Answer, error) {
	answers, err := answerService.answerRepo.GetAnswerByQuestionID(ctx, nil, questionID)
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (answerService *answerService) GetAnswerByStudentID(ctx context.Context, id dto.GetAnswerByStudentIDRequest) ([]entities.Answer, error) {
	answers, err := answerService.answerRepo.GetAnswerByStudentID(ctx, nil, id)
	if err != nil {
		return nil, err
	}
	return answers, nil
}