package dto

import (
	entities "assesment/entities"
	"time"

	"github.com/google/uuid"
)

type CreateAllQuestionRequest struct{
	QuestionText string `json:"question_text" binding:"required"`
	EvaluationID uuid.UUID `gorm:"type:uuid" json:"evaluation_id" binding:"required"`
	Choices []ChoiceCreateRequest `json:"choices"`
}

// type CreateAllQuestionRequest struct{
// 	EvaluationID uuid.UUID `gorm:"type:uuid" json:"evaluation_id"`
// 	Data []CreateAllQuestionData `json:"data"`
// }

type UpdateAllQuestionRequest struct{
	QuestionText string `json:"question_text"`
	EvaluationID uuid.UUID `gorm:"type:uuid" json:"evaluation_id"`
	Choices []ChoiceUpdateRequest `json:"choices"`
}

type QuestionCreateRequest struct {
	QuestionText string `json:"question_text"`
	EvaluationID uuid.UUID `gorm:"type:uuid" json:"evaluation_id"`
}

type QuestionUpdateRequest struct {
	Id           string `json:"id"`
	QuestionText string `json:"question_text"`
	EvaluationID uuid.UUID `gorm:"type:uuid" json:"evaluation_id"`
}

type QuestionResponse struct {
	ID           uuid.UUID `json:"id"`
	QuestionText string `json:"question_text"`
	EvaluationID uuid.UUID `gorm:"type:uuid" json:"evaluation_id"`
	CreatedAt   time.Time `json:"created_at"`
	ChoiceResponse []ChoiceResponse `json:"choices"`
}

type GetAllQuestionsResponse struct {
	Questions []entities.Question `json:"questions"`
}

