package dto

import (
	entities "assesment/entities"

	"github.com/google/uuid"
)

type SubmissionCreateRequest struct {
	UserID	   uuid.UUID `json:"user_id" binding:"required"`
	AssessmentID uuid.UUID `json:"assessment_id" binding:"required"`
}

type SubmissionCreateResponse struct {
	ID           uuid.UUID      `gorm:"type:uuid" json:"id"`
	UserID       uuid.UUID      `gorm:"type:uuid" json:"user_id"`
	AssessmentID uuid.UUID      `gorm:"type:uuid" json:"assessment_id"`
	Question []entities.Question `json:"question"`
}
