package dto

import (
	entities "assesment/entities"
	"time"

	"github.com/google/uuid"
)

type CreateAllQuestionRequest struct {
	AssessmentID uuid.UUID `gorm:"type:uuid" json:"assessment_id" binding:"required"`
	Questions []AllQuestionRequest `json:"questions" binding:"required,dive"`
}
type AllQuestionRequest struct{
	QuestionText string `json:"question_text" binding:"required"`
	Choices []ChoiceCreateRequest `json:"choices" binding:"required,dive"`
}

type QuestionCreateRequest struct {
	QuestionText string `json:"question_text"`
	AssessmentID uuid.UUID `gorm:"type:uuid" json:"assessment_id"`
}


type QuestionUpdateRequest struct {
	QuestionId           uuid.UUID `json:"question_id" binding:"required"`
	QuestionText string `json:"question_text"`
	Choices      []ChoiceCreateRequest `json:"choices" binding:"required,dive"`
}

type QuestionResponse struct {
	AssessmentID uuid.UUID `gorm:"type:uuid" json:"assessment_id"`
	ID           uuid.UUID `json:"question_id"`
	QuestionText string `json:"question_text"`
	CreatedAt   time.Time `json:"created_at"`
	ChoiceResponse []ChoiceResponse `json:"choices"`
}

type GetAllQuestionsResponse struct {
	Questions []entities.Question `json:"questions"`
}

