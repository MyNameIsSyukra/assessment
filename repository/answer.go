package repository

import (
	"assesment/dto"
	entities "assesment/entities"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	AnswerRepository interface {
		CreateAnswer(ctx context.Context, tx *gorm.DB, answer *entities.Answer) (dto.AnswerResponse, error)
		GetAnswerByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*entities.Answer, error)
		GetAllAnswers() ([]entities.Answer, error)
		UpdateAnswer(ctx context.Context, tx *gorm.DB, answer *entities.Answer) (*entities.Answer, error)
		GetAnswerByQuestionID(ctx context.Context, tx *gorm.DB, questionID uuid.UUID) ([]entities.Answer, error)
		// GetAnswerByStudentID(ctx context.Context, tx *gorm.DB, id dto.GetAnswerByStudentIDRequest) ([]entities.Answer, error)
		GetAnswerBySubmissionID(ctx context.Context, tx *gorm.DB, submissionID uuid.UUID) ([]dto.GetAnswerBySubmissionIDResponse, error)
		GetAnswerBySubmissionIDAndQuestionID(ctx context.Context, tx *gorm.DB, SubmisiionID uuid.UUID, IdQuestion uuid.UUID) (*entities.Answer, error)
	}
	answerRepository struct {
		Db *gorm.DB
	}
)

func NewAnswerRepository(db *gorm.DB) AnswerRepository {
	return &answerRepository{Db: db}
}

func (answerRepo *answerRepository) CreateAnswer(ctx context.Context, tx *gorm.DB, answer *entities.Answer) (dto.AnswerResponse, error) {
	if err := answerRepo.Db.Create(answer).Error; err != nil {
		return dto.AnswerResponse{}, err
	}


	res := dto.AnswerResponse{
		ID:        answer.ID,
		IdQuestion: answer.QuestionID,
		SubmissionID: answer.SubmissionID, 
		IdChoice:  answer.ChoiceID,
		CreatedAt: answer.CreatedAt,
	}
	return res, nil
}

func (answerRepo *answerRepository) GetAnswerByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*entities.Answer, error) {
	var answer entities.Answer
	if err := answerRepo.Db.Where("id = ?", id).Preload("Submission").First(&answer).Error; err != nil {
		return &entities.Answer{}, err
	}
	return &answer, nil
}

func (answerRepo *answerRepository) GetAllAnswers() ([]entities.Answer, error) {
	var answers []entities.Answer
	if err := answerRepo.Db.Find(&answers).Error; err != nil {
		return []entities.Answer{}, err
	}
	return answers, nil
}

func (answerRepo *answerRepository) UpdateAnswer(ctx context.Context, tx *gorm.DB, answer *entities.Answer) (*entities.Answer, error) {
	if err := answerRepo.Db.Where("id = ?",answer.ID).Updates(answer).Error; err != nil {
		return &entities.Answer{}, err
	}
	return answer, nil
}

func (answerRepo *answerRepository) GetAnswerByQuestionID(ctx context.Context, tx *gorm.DB, questionID uuid.UUID) ([]entities.Answer, error) {
	var answers []entities.Answer
	if err := answerRepo.Db.Where("question_id = ?", questionID).Find(&answers).Error; err != nil {
		return []entities.Answer{}, err
	}
	return answers, nil
}

func (answerRepo *answerRepository) GetAnswerBySubmissionID(ctx context.Context, tx *gorm.DB, submissionID uuid.UUID) ([]dto.GetAnswerBySubmissionIDResponse, error) {
	var answers []entities.Answer
	var dtoAnswers []dto.GetAnswerBySubmissionIDResponse
	if err := answerRepo.Db.Where("submission_id = ?", submissionID).Preload("Question").Preload("Choice").Find(&answers).Error; err != nil {
		return []dto.GetAnswerBySubmissionIDResponse{}, err
	}
	for _, answer := range answers {
		choice := dto.ChoiceGetAnswerBySubmissionIDResponse{
			ID:         answer.Choice.ID,
			ChoiceText: answer.Choice.ChoiceText,
			QuestionID: answer.Choice.QuestionID,
			CreatedAt:  answer.Choice.CreatedAt,
			UpdatedAt:  answer.Choice.UpdatedAt,
			DeletedAt:  answer.Choice.DeletedAt,
		}
		dtoAnswer := dto.GetAnswerBySubmissionIDResponse{
			ID:         answer.ID,
			SubmissionID: answer.SubmissionID,
			IdQuestion: answer.QuestionID,
			IdChoice:   answer.ChoiceID,
			CreatedAt:  answer.CreatedAt,
			Question:   answer.Question,
			Choice:     choice,
		}
		dtoAnswers = append(dtoAnswers, dtoAnswer)
	}

	return dtoAnswers, nil
}

func (answerRepo *answerRepository) GetAnswerBySubmissionIDAndQuestionID(ctx context.Context, tx *gorm.DB, SubmisiionID uuid.UUID, IdQuestion uuid.UUID) (*entities.Answer, error) {
	var answers entities.Answer
	if err := answerRepo.Db.Where("submission_id = ? AND question_id = ?", SubmisiionID, IdQuestion).Find(&answers).Error; err != nil {
		return &entities.Answer{}, err
	}
	return &answers, nil
}