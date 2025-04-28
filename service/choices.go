package service

import (
	dto "assesment/dto"
	entities "assesment/entities"
	repository "assesment/repository"
	"context"
)

type (
	ChoiceService interface {
		CreateChoice(ctx context.Context, choice *dto.ChoiceCreateRequest) (dto.ChoiceResponse, error)
		GetChoiceByID(ctx context.Context, id string) (dto.ChoiceResponse, error)
		UpdateChoice(ctx context.Context, choice *dto.ChoiceUpdateRequest) (*entities.Choice, error)
		GetChoiceByQuestionID(ctx context.Context, questionID string) ([]entities.Choice, error)
	}
	choiceService struct {
		choiceRepo repository.ChoiceRepository
	}
)

func NewChoiceService(choiceRepo repository.ChoiceRepository) ChoiceService {
	return &choiceService{
		choiceRepo: choiceRepo,
	}
}

func (choiceService *choiceService) CreateChoice(ctx context.Context, choice *dto.ChoiceCreateRequest) (dto.ChoiceResponse, error) {
		choiceEntity := entities.Choice{
			ChoiceText: choice.ChoiceText,
			IsCorrect:  choice.IsCorrect,
			// QuestionID: choice.QuestionID,
		}

		createdChoice, err := choiceService.choiceRepo.CreateChoice(ctx, nil, &choiceEntity)
		if err != nil {
			return dto.ChoiceResponse{}, err
		}
		return dto.ChoiceResponse{
			ID:        createdChoice.ID,
			ChoiceText: createdChoice.ChoiceText,
			IsCorrect:  createdChoice.IsCorrect,
			QuestionID: createdChoice.QuestionID,
		}, nil
}

func (choiceService *choiceService) GetChoiceByID(ctx context.Context, id string) (dto.ChoiceResponse, error) {
	choice, err := choiceService.choiceRepo.GetChoiceByID(ctx, nil, id)
	if err != nil {
		return dto.ChoiceResponse{}, err
	}
	return choice, nil
}

func (choiceService *choiceService) UpdateChoice(ctx context.Context, choice *dto.ChoiceUpdateRequest) (*entities.Choice, error) {
	choiceEntity, err := choiceService.choiceRepo.GetChoiceByID(ctx, nil, choice.ID)
	if err != nil {
		return nil, err
	}
	
	data := entities.Choice{
		ID: choiceEntity.ID,
		ChoiceText: choice.ChoiceText,
		IsCorrect: choice.IsCorrect,		
		QuestionID: choiceEntity.QuestionID,
	}		
	updatedChoice, err := choiceService.choiceRepo.UpdateChoice(ctx, nil, &data)
	
	if err != nil {
		return nil, err
	}
	return updatedChoice, nil
}

func (choiceService *choiceService) GetChoiceByQuestionID(ctx context.Context, questionID string) ([]entities.Choice, error) {
	choices, err := choiceService.choiceRepo.GetChoiceByQuestionID(ctx, nil, questionID)
	if err != nil {
		return nil, err
	}
	return choices, nil
}

