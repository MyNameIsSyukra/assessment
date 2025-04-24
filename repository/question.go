package repository

import (
	entities "assesment/entities"
	"context"

	"gorm.io/gorm"
)

type (
	QuestionRepository interface {
		CreateQuestion(ctx context.Context, tx *gorm.DB, question *entities.Question) (*entities.Question, error)
		GetQuestionByID(ctx context.Context, tx *gorm.DB, id string) (*entities.Question, error)
		GetAllQuestions() ([]entities.Question, error)
		UpdateQuestion(ctx context.Context, tx *gorm.DB, question *entities.Question) (*entities.Question, error)
		DeleteQuestion(ctx context.Context, tx *gorm.DB, id string) error
		GetQuestionsByAssessmentID(ctx context.Context, tx *gorm.DB, assessmentID string) ([]entities.Question, error)
	}
	questionRepository struct {
		Db *gorm.DB
	}
)

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepository{Db: db}
}

func (questionRepo *questionRepository) CreateQuestion(ctx context.Context, tx *gorm.DB, question *entities.Question) (*entities.Question, error) {	
	if err := questionRepo.Db.Create(question).Error; err != nil {
		return nil, err
	}
	return question, nil
}

func (questionRepo *questionRepository) GetQuestionByID(ctx context.Context, tx *gorm.DB, id string) (*entities.Question, error) {
	var question entities.Question
	if err := questionRepo.Db.Where("id = ?", id).First(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func (questionRepo *questionRepository) GetAllQuestions() ([]entities.Question, error) {
	var questions []entities.Question
	if err := questionRepo.Db.Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func (questionRepo *questionRepository) UpdateQuestion(ctx context.Context, tx *gorm.DB, question *entities.Question) (*entities.Question, error) {
	if err := questionRepo.Db.Save(question).Error; err != nil {
		return nil, err
	}
	return question, nil
}

func (questionRepo *questionRepository) DeleteQuestion(ctx context.Context, tx *gorm.DB, id string) error {
	var question entities.Question
	if err := questionRepo.Db.Delete(&question, id).Error; err != nil {
		return err
	}
	return nil
}

func (questionRepo *questionRepository) GetQuestionsByAssessmentID(ctx context.Context, tx *gorm.DB, assessmentID string) ([]entities.Question, error) {
	var questions []entities.Question
	if err := questionRepo.Db.Where("assessment_id = ?", assessmentID).Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

