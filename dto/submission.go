package dto

import (
	"github.com/google/uuid"
)

type SubmissionCreateRequest struct {
	UserID	   uuid.UUID `json:"user_id"`
	AssessmentID uuid.UUID `json:"assessment_id"`
}

type SubmissionCreateResponse struct {
	ID           uuid.UUID      `gorm:"type:uuid" json:"id"`
	UserID       uuid.UUID      `gorm:"type:uuid" json:"user_id"`
	AssessmentID uuid.UUID      `gorm:"type:uuid" json:"assessment_id"`
}
