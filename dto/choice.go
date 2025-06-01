package dto

import (
	"github.com/google/uuid"
)

type ChoiceCreateRequest struct {
	ChoiceText string    `json:"choice_text" binding:"required"`
	IsCorrect  bool      `json:"is_correct"`
	QuestionID uuid.UUID `gorm:"type:uuid" json:"question_id"`
}

type ChoiceUpdateRequest struct {
	ID        uuid.UUID `json:"choice_id" binding:"required"`
	ChoiceText string `json:"choice_text" binding:"required"`
	IsCorrect  bool   `json:"is_correct"`
	QuestionID uuid.UUID `gorm:"type:uuid" json:"question_id" binding:"required"`
}

type ChoiceResponse struct {
	ID        uuid.UUID `json:"id"`
	ChoiceText string    `json:"choice_text"`
	IsCorrect  bool      `json:"is_correct"`
	QuestionID uuid.UUID `gorm:"type:uuid" json:"question_id"`
}