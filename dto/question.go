package dto

import "github.com/google/uuid"

type QuestionCreateRequest struct {
	QuestionText string `json:"question_text"`
	EvaluationID uuid.UUID `gorm:"type:uuid" json:"evaluation_id"`
}

type QuestionUpdateRequest struct {
	Id           string `json:"id"`
	QuestionText string `json:"question_text"`
	EvaluationID uuid.UUID `gorm:"type:uuid" json:"evaluation_id"`
}
