package service

import (
	dto "assesment/dto"
	entities "assesment/entities"
	repository "assesment/repository"
	"context"
	"fmt"
)

type (
	QuestionService interface {
		CreateQuestion(ctx context.Context, question *dto.QuestionCreateRequest) (*entities.Question, error)
		GetAllQuestions(ctx context.Context) ([]entities.Question, error)
		GetQuestionByID(ctx context.Context, id string) (*entities.Question, error)
		UpdateQuestion(ctx context.Context, question *dto.QuestionUpdateRequest) (*entities.Question, error)
		DeleteQuestion(ctx context.Context, id string) error
		GetQuestionsByAssessmentID(ctx context.Context, assessmentID string) ([]entities.Question, error)
	}
	questionService struct {
		questionRepo repository.QuestionRepository
	}
)

func NewQuestionService(questionRepo repository.QuestionRepository) QuestionService {
	return &questionService{
		questionRepo: questionRepo,
	}
}

func (questionService *questionService) CreateQuestion(ctx context.Context, question *dto.QuestionCreateRequest) (*entities.Question, error) {
		questionEntity := entities.Question{
			QuestionText: question.QuestionText,
			EvaluationId: question.EvaluationId,
		}
		
		createdQuestion, err := questionService.questionRepo.CreateQuestion(ctx, nil, &questionEntity)
		if err != nil {
			return nil, err
		}
		return createdQuestion, nil	
}

func (questionService *questionService) GetAllQuestions(ctx context.Context) ([]entities.Question, error) {
	questions, err := questionService.questionRepo.GetAllQuestions()
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (questionService *questionService) GetQuestionByID(ctx context.Context, id string) (*entities.Question, error) {
	question, err := questionService.questionRepo.GetQuestionByID(ctx, nil, id)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (questionService *questionService) UpdateQuestion(ctx context.Context, question *dto.QuestionUpdateRequest) (*entities.Question, error) {
	
	questionEntity,err := questionService.questionRepo.GetQuestionByID(ctx, nil, question.Id)
	if questionEntity == nil {
		return nil, fmt.Errorf("question not found")
	}
	data := entities.Question{
		QuestionText: question.QuestionText,
		EvaluationId: question.EvaluationId,
	}
	updatedQuestion, err := questionService.questionRepo.UpdateQuestion(ctx, nil, &data)
	if err != nil {
		return nil, err
	}
	return updatedQuestion, nil
}

func (questionService *questionService) DeleteQuestion(ctx context.Context, id string) error {
	question, err := questionService.questionRepo.GetQuestionByID(ctx, nil, id)
	if question == nil {
		return fmt.Errorf("question not found")
	}
	err = questionService.questionRepo.DeleteQuestion(ctx, nil, id)
	if err != nil {
		return err
	}
	return nil
}

func (questionService *questionService) GetQuestionsByAssessmentID(ctx context.Context, assessmentID string) ([]entities.Question, error) {
	questions, err := questionService.questionRepo.GetQuestionsByAssessmentID(ctx, nil, assessmentID)
	if err != nil {
		return nil, err
	}
	return questions, nil
}