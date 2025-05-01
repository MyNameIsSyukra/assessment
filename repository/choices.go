package repository

import (
	"assesment/dto"
	entities "assesment/entities"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	ChoiceRepository interface {
		CreateChoice(ctx context.Context, tx *gorm.DB, choice *entities.Choice) (*entities.Choice, error)
		GetChoiceByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (dto.ChoiceResponse, error)
		UpdateChoice(ctx context.Context, tx *gorm.DB, choice *entities.Choice) (*entities.Choice, error)
		GetChoiceByQuestionID(ctx context.Context, tx *gorm.DB, questionID uuid.UUID) ([]entities.Choice, error)
	}
	choiceRepository struct {
		Db *gorm.DB
	}
)

func NewChoiceRepository(db *gorm.DB) ChoiceRepository {
	return &choiceRepository{Db: db}
}

func (choiceRepo *choiceRepository) CreateChoice(ctx context.Context, tx *gorm.DB, choice *entities.Choice) (*entities.Choice, error) {
	if err := choiceRepo.Db.Create(choice).Error; err != nil {
		return nil, err
	}
	return choice, nil
}

func (choiceRepo *choiceRepository) GetChoiceByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (dto.ChoiceResponse, error) {
	var choice entities.Choice
	if err := choiceRepo.Db.Where("id = ?", id).First(&choice).Error; err != nil {
		return dto.ChoiceResponse{}, err
	}

	data := dto.ChoiceResponse{
		ID: choice.ID,
		ChoiceText: choice.ChoiceText,
		IsCorrect: choice.IsCorrect,
		QuestionID: choice.QuestionID,
	}
	return data, nil
}

func (choiceRepo *choiceRepository) UpdateChoice(ctx context.Context, tx *gorm.DB, choice *entities.Choice) (*entities.Choice, error) {
	if err := choiceRepo.Db.Where("id = ?", choice.ID).Updates(choice).Error; err != nil {
		return nil, err
	}
	return choice, nil
}

func (choiceRepo *choiceRepository) GetChoiceByQuestionID(ctx context.Context, tx *gorm.DB, questionID uuid.UUID) ([]entities.Choice, error) {
	var choices []entities.Choice
	if err := choiceRepo.Db.Where("question_id = ?", questionID).Find(&choices).Error; err != nil {
		return nil, err
	}
	return choices, nil
}

