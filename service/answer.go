package service

import (
	dto "assesment/dto"
	entities "assesment/entities"
	repository "assesment/repository"
	"context"
)

type (
	AnswerService interface {
		CreateAnswer(ctx context.Context, answer *dto.AnswerCreateRequest) (*entities.Answer, error)
		GetAllAnswers(ctx context.Context) ([]entities.Answer, error)
		GetAnswerByID(ctx context.Context, id string) (*entities.Answer, error)
		UpdateAnswer(ctx context.Context, answer *dto.AnswerUpdateRequest) (*entities.Answer, error)
		GetAnswerByQuestionID(ctx context.Context, questionID string) ([]entities.Answer, error)
		GetAnswerByStudentID(ctx context.Context, userID string) ([]entities.Answer, error)
	}
	answerService struct {
		answerRepo repository.AnswerRepository
	}
)