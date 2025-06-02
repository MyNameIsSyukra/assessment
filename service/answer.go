package service

import (
	dto "assesment/dto"
	entities "assesment/entities"
	repository "assesment/repository"
	"assesment/utils"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	AnswerService interface {
		CreateAnswer(ctx context.Context, answer *dto.AnswerCreateRequest) (dto.AnswerResponse, error)
		GetAllAnswers(ctx context.Context) ([]entities.Answer, error)
		GetAnswerByID(ctx context.Context, id uuid.UUID) (*entities.Answer, error)
		UpdateAnswer(ctx context.Context, answer *dto.AnswerUpdateRequest) (*dto.AnswerUpdatedResponse, error)
		GetAnswerByQuestionID(ctx context.Context, questionID uuid.UUID) ([]entities.Answer, error)
		ContinueSubmission(ctx context.Context, submissionID uuid.UUID) (*dto.ContinueSubmissionIDResponse, error)
		// GetAnswerByStudentID(ctx context.Context, id dto.GetAnswerByStudentIDRequest) ([]entities.Answer, error)
	}
	answerService struct {
		answerRepo     repository.AnswerRepository
		submissionRepo repository.SubmissionRepository
		assesmentRepo repository.AssessmentRepository
		questionRepo repository.QuestionRepository
	}
)

func NewAnswerService(answerRepo repository.AnswerRepository,submissionRepo repository.SubmissionRepository, assesmentRepo repository.AssessmentRepository, questionRepo repository.QuestionRepository) AnswerService {
	return &answerService{
		answerRepo: answerRepo,
		submissionRepo: submissionRepo,
		assesmentRepo: assesmentRepo,
		questionRepo: questionRepo,
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

func (answerService *answerService) ContinueSubmission(ctx context.Context, submissionID uuid.UUID) (*dto.ContinueSubmissionIDResponse, error) {
	// Validate input
	if submissionID == uuid.Nil {
		return nil, errors.New("invalid submission ID")
	}

	// Get submission
	submission, err := answerService.submissionRepo.GetSubmissionByID(ctx, nil, submissionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("submission not found")
		}
		return nil, fmt.Errorf("failed to get submission: %w", err)
	}

	// Check submission status
	if submission.Status != "in_progress" {
		return nil, errors.New("submission is not in progress")
	}

	// Get answers for this submission
	answers, err := answerService.answerRepo.GetAnswerBySubmissionID(ctx, nil, submissionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get answers: %w", err)
	}

	// Get questions for the assessment
	questions, err := answerService.questionRepo.GetQuestionsByAssessmentID(ctx, nil, submission.AssessmentID)
	if err != nil {
		return nil, fmt.Errorf("failed to get questions: %w", err)
	}

	// Create answer lookup map for better performance
	answerMap := make(map[uuid.UUID]*entities.Answer)
	for i := range answers {
		answerMap[answers[i].QuestionID] = &answers[i]
	}

	// Build response
	questionsSubmittedAnswer := make([]dto.QuestionsSubmittedAnswer, 0, len(questions))
	
	for _, q := range questions {
		questionAnswer := dto.QuestionsSubmittedAnswer{
			QuestionID:   q.ID,
			QuestionText: q.QuestionText,
			AssessmentID: q.AssessmentID,
			CreatedAt:    q.CreatedAt,
			UpdatedAt:    q.UpdatedAt,
			DeletedAt:    q.DeletedAt,
			Choice:       q.Choices, // Assuming Choices is a slice of entities.Choice
		}

		// Check if answer exists for this question
		if answerFound, exists := answerMap[q.ID]; exists {
			// Handle potential nil pointer issue

			questionAnswer.Answers = &dto.Answer{
				ID:           answerFound.ID,
				QuestionID:   answerFound.QuestionID,
				ChoiceID:     answerFound.ChoiceID,
				ChoiceText:   answerFound.Choice.ChoiceText,
				SubmissionID: answerFound.SubmissionID,
				CreatedAt:    answerFound.CreatedAt,
				UpdatedAt:    answerFound.UpdatedAt,
				DeletedAt:    answerFound.DeletedAt,
			}
		}
		// If no answer found, Answers will remain nil

		questionsSubmittedAnswer = append(questionsSubmittedAnswer, questionAnswer)
	}

	return &dto.ContinueSubmissionIDResponse{
		Assessment_ID: submission.AssessmentID,
		User_ID:       submission.UserID,
		SubmissionID:  submissionID,
		EndedTime:     submission.EndedTime,
		Question:      questionsSubmittedAnswer,
	}, nil
}