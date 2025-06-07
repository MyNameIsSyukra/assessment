package service

import (
	dto "assesment/dto"
	entities "assesment/entities"
	repository "assesment/repository"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	QuestionService interface {
		CreateQuestion(ctx context.Context, question *dto.QuestionCreateRequest) (dto.QuestionResponse, error)
		// GetAllQuestions(ctx context.Context) (dto.GetAllQuestionsResponse, error)
		GetQuestionByID(ctx context.Context, id uuid.UUID) (*entities.Question, error)
		UpdateQuestion(ctx context.Context, question *dto.QuestionUpdateRequest) (*entities.Question, error)
		DeleteQuestion(ctx context.Context, id uuid.UUID) error
		GetQuestionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]entities.Question, error)
		CreatePertanyaan(ctx context.Context, AssessmentID uuid.UUID, question dto.AllQuestionRequest) (dto.QuestionResponse, error)
	}
	questionService struct {
		questionRepo repository.QuestionRepository
		assesRepo  repository.AssessmentRepository
		choiceRepo repository.ChoiceRepository
	}
)

func NewQuestionService(questionRepo repository.QuestionRepository, assesRepo repository.AssessmentRepository, choiceRepo repository.ChoiceRepository) QuestionService {
	return &questionService{
		questionRepo: questionRepo,
		assesRepo:  assesRepo,
		choiceRepo: choiceRepo,
	}
}

func (questionService *questionService) CreatePertanyaan(ctx context.Context, AssessmentID uuid.UUID, question dto.AllQuestionRequest) (dto.QuestionResponse, error) {
	// Check if the assessment exists
	assesment, err := questionService.assesRepo.GetAssessmentByID(ctx, nil, AssessmentID)
	if assesment == nil {
		return dto.QuestionResponse{}, errors.New("assessment not found")
	}
	if err != nil {
		return dto.QuestionResponse{}, err
	}
	questionEntity := entities.Question{
		AssessmentID: AssessmentID,
		QuestionText: question.QuestionText,
	}
	if time.Now().After(assesment.StartTime) || assesment.EndTime.Before(time.Now()) {
		return dto.QuestionResponse{}, errors.New("assessment is already started or ended")
	}
	createdQuestion, err := questionService.questionRepo.CreateQuestion(ctx, nil, &questionEntity)
	if err != nil {
		return dto.QuestionResponse{}, err
	}
	
	for _, choice := range question.Choices {
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
			AssessmentID: question.AssessmentID,
		}
		
		createdQuestion, err := questionService.questionRepo.CreateQuestion(ctx, nil, &questionEntity)
		if err != nil {
			return dto.QuestionResponse{}, err
		}
		return createdQuestion, nil	
}

func (questionService *questionService) GetQuestionByID(ctx context.Context, id uuid.UUID) (*entities.Question, error) {
	question, err := questionService.questionRepo.GetQuestionByID(ctx, nil, id)
	if err != nil {
		return &entities.Question{}, err
	}
	return question, nil
}

func (questionService *questionService) UpdateQuestion(ctx context.Context, question *dto.QuestionUpdateRequest) (*entities.Question, error) {
	questionEntity, err := questionService.questionRepo.GetQuestionByID(ctx, nil, question.QuestionId)
	// fmt.Println(questionEntity
	if err != nil {
		return &entities.Question{}, err
	}
	// fmt.Println("aftercheck entity" )
	data := entities.Question{
		ID: questionEntity.ID,
		QuestionText: question.QuestionText,
	}
	assesment, err := questionService.assesRepo.GetAssessmentByID(ctx, nil, questionEntity.AssessmentID)
	if err != nil {
		return &entities.Question{}, err
	}
	if assesment.StartTime.After(time.Now()) || assesment.EndTime.Before(time.Now()) {
		return &entities.Question{}, errors.New("cannot edit question, assessment is already started or ended")
	}
	// fmt.Println(question)
	_, err = questionService.questionRepo.UpdateQuestion(ctx, nil, &data)
	if err != nil {
		return &entities.Question{}, err
	}

	err = questionService.choiceRepo.DeleteChoiceByQuestionID(ctx, nil, question.QuestionId)
	if err != nil {
		return nil, err
	}
	for _, choice := range question.Choices {
		data := entities.Choice{
			ChoiceText: choice.ChoiceText,
			IsCorrect:  choice.IsCorrect,
			QuestionID: questionEntity.ID,
		}
		_, err = questionService.choiceRepo.CreateChoice(ctx, nil, &data)
		if err != nil {
			return nil, err
		}

	}
	afterQuestion, err := questionService.questionRepo.GetQuestionByID(ctx, nil, questionEntity.ID)
	if err != nil {
		return &entities.Question{}, err
	}
	return afterQuestion, nil
}

func (questionService *questionService) DeleteQuestion(ctx context.Context, id uuid.UUID) error {
	question, err := questionService.questionRepo.GetQuestionByID(ctx, nil, id)
	if err != nil {
		return err
	}
	
	err = questionService.questionRepo.DeleteQuestion(ctx, nil, question.ID)
	if err != nil {
		return err
	}
	return nil
}

func (questionService *questionService) GetQuestionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]entities.Question, error) {
	questions, err := questionService.questionRepo.GetQuestionsByAssessmentID(ctx, nil, assessmentID)
	if err != nil {
		return []entities.Question{}, err
	}
	if len(questions) == 0 {
		return []entities.Question{}, err
	}
	return questions, nil
}


// func (questionService *questionService) GetAllQuestions(ctx context.Context) (dto.GetAllQuestionsResponse, error) {
// 	questions, err := questionService.questionRepo.GetAllQuestions()
// 	if err != nil {
// 		return dto.GetAllQuestionsResponse{}, err
// 	}
// 	return questions, nil
// }