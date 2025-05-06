package dto

import (
	entities "assesment/entities"
	"time"

	"github.com/google/uuid"
)

type CreateAllQuestionRequest struct {
	EvaluationID uuid.UUID `gorm:"type:uuid" json:"evaluation_id" binding:"required"`
	Questions []AllQuestionRequest `json:"questions" binding:"required,dive"`
}
type AllQuestionRequest struct{
	QuestionText string `json:"question_text" binding:"required"`
	Choices []ChoiceCreateRequest `json:"choices" binding:"required,dive"`
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
	Id           uuid.UUID `json:"id" binding:"required"`
	QuestionText string `json:"question_text"`
	EvaluationID uuid.UUID `gorm:"type:uuid" json:"evaluation_id" binding:"required"`
	Choices      []ChoiceCreateRequest `json:"choices" binding:"required,dive"`
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

