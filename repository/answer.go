package repository

import (
	entities "assesment/entities"
	"context"

	"gorm.io/gorm"
)

type (
	AnswerRepository interface {
		CreateAnswer(ctx context.Context, tx *gorm.DB, answer *entities.Answer) (*entities.Answer, error)
		GetAnswerByID(ctx context.Context, tx *gorm.DB, id string) (*entities.Answer, error)
		GetAllAnswers() ([]entities.Answer, error)
		UpdateAnswer(ctx context.Context, tx *gorm.DB, answer *entities.Answer) (*entities.Answer, error)
		GetAnswerByQuestionID(ctx context.Context, tx *gorm.DB, questionID string) ([]entities.Answer, error)
		GetAnswerByStudentID(ctx context.Context, tx *gorm.DB, userID string) ([]entities.Answer, error)
	}
	answerRepository struct {
		Db *gorm.DB
	}
)

func NewAnswerRepository(db *gorm.DB) AnswerRepository {
	return &answerRepository{Db: db}
}

func (answerRepo *answerRepository) CreateAnswer(ctx context.Context, tx *gorm.DB, answer *entities.Answer) (*entities.Answer, error) {
	if err := answerRepo.Db.Create(answer).Error; err != nil {
		return nil, err
	}
	return answer, nil
}

func (answerRepo *answerRepository) GetAnswerByID(ctx context.Context, tx *gorm.DB, id string) (*entities.Answer, error) {
	var answer entities.Answer
	if err := answerRepo.Db.Where("id = ?", id).First(&answer).Error; err != nil {
		return nil, err
	}
	return &answer, nil
}

func (answerRepo *answerRepository) GetAllAnswers() ([]entities.Answer, error) {
	var answers []entities.Answer
	if err := answerRepo.Db.Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil
}

func (answerRepo *answerRepository) UpdateAnswer(ctx context.Context, tx *gorm.DB, answer *entities.Answer) (*entities.Answer, error) {
	if err := answerRepo.Db.Save(answer).Error; err != nil {
		return nil, err
	}
	return answer, nil
}

func (answerRepo *answerRepository) GetAnswerByQuestionID(ctx context.Context, tx *gorm.DB, questionID string) ([]entities.Answer, error) {
	var answers []entities.Answer
	if err := answerRepo.Db.Where("question_id = ?", questionID).Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil
}

func (answerRepo *answerRepository) GetAnswerByStudentID(ctx context.Context, tx *gorm.DB, userID string) ([]entities.Answer, error) {
	var answers []entities.Answer
	if err := answerRepo.Db.Where("st = ?", userID).Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil
}
