package service

import (
	dto "assesment/dto"
	entities "assesment/entities"
	repository "assesment/repository"
	"context"
)

type (
	QuestionService interface {
		CreateQuestion(ctx context.Context, question *dto.QuestionCreateRequest) (dto.QuestionResponse, error)
		GetAllQuestions(ctx context.Context) (dto.GetAllQuestionsResponse, error)
		GetQuestionByID(ctx context.Context, id string) (*entities.Question, error)
		UpdateQuestion(ctx context.Context, question *dto.QuestionUpdateRequest) (*entities.Question, error)
		DeleteQuestion(ctx context.Context, id string) error
		GetQuestionsByAssessmentID(ctx context.Context, assessmentID string) ([]entities.Question, error)
		CreatePertanyaan(ctx context.Context, req dto.CreateAllQuestionRequest) (dto.QuestionResponse, error)
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

func (questionService *questionService) CreatePertanyaan(ctx context.Context,req dto.CreateAllQuestionRequest) (dto.QuestionResponse, error) {
	questionEntity := entities.Question{
		QuestionText: req.QuestionText,
		EvaluationID: req.EvaluationID,
	}
	createdQuestion, err := questionService.questionRepo.CreateQuestion(ctx, nil, &questionEntity)
	if err != nil {
		return dto.QuestionResponse{}, err
	}
	
	for _, choice := range req.Choices {
		choice.QuestionID = createdQuestion.ID
		data := entities.Choice{
			ChoiceText: choice.ChoiceText,
			IsCorrect:  choice.IsCorrect,
			QuestionID: choice.QuestionID,
		}
		datas, err := questionService.questionRepo.CreateChoice(ctx, nil, &data)
		if err != nil {
			return dto.QuestionResponse{}, err
		}
		createdQuestion.ChoiceResponse = append(createdQuestion.ChoiceResponse, dto.ChoiceResponse{
			ID: datas.ID,
			ChoiceText: choice.ChoiceText,
			IsCorrect: choice.IsCorrect,
			QuestionID: choice.QuestionID,
		})
	}
	
	return createdQuestion, nil
}

func (questionService *questionService) CreateQuestion(ctx context.Context, question *dto.QuestionCreateRequest) (dto.QuestionResponse, error) {
		questionEntity := entities.Question{
			QuestionText: question.QuestionText,
			EvaluationID: question.EvaluationID,
		}
		
		createdQuestion, err := questionService.questionRepo.CreateQuestion(ctx, nil, &questionEntity)
		if err != nil {
			return dto.QuestionResponse{}, err
		}
		return createdQuestion, nil	
}

func (questionService *questionService) GetAllQuestions(ctx context.Context) (dto.GetAllQuestionsResponse, error) {
	questions, err := questionService.questionRepo.GetAllQuestions()
	if err != nil {
		return dto.GetAllQuestionsResponse{}, err
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
	questionEntity, err := questionService.questionRepo.GetQuestionByID(ctx, nil, question.Id)
	// fmt.Println(questionEntity
	if err != nil {
		return nil, err
	}
	data := entities.Question{
		ID: questionEntity.ID,
		QuestionText: question.QuestionText,
		EvaluationID: question.EvaluationID,
	}
	// fmt.Println(question)
	updatedQuestion, err := questionService.questionRepo.UpdateQuestion(ctx, nil, &data)
	if err != nil {
		return nil, err
	}
	return updatedQuestion, nil
}

func (questionService *questionService) DeleteQuestion(ctx context.Context, id string) error {
	question, err := questionService.questionRepo.GetQuestionByID(ctx, nil, id)
	if err != nil {
		return err
	}
	
	err = questionService.questionRepo.DeleteQuestion(ctx, nil, question.ID.String())
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

