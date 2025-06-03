package dto

import (
	"assesment/entities"
	"time"

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

type PublicQuestionResponse struct {
    ID           uuid.UUID       `json:"question_id"`
    QuestionText string          `json:"question_text"`
    AssessmentID uuid.UUID       `json:"assessment_id"`
    CreatedAt    time.Time       `json:"created_at"`
    UpdatedAt    time.Time       `json:"updated_at"`
    Choices      []PublicChoiceResponse `json:"choice"`
}

type PublicChoiceResponse struct {
    ID         uuid.UUID `json:"choice_id"`
    ChoiceText string    `json:"choice_text"`
    QuestionID uuid.UUID `json:"question_id"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
    // IsCorrect sengaja tidak disertakan untuk public response
}

// Converter functions
func ToQuestionResponse(question entities.Question) PublicQuestionResponse {
    choices := make([]PublicChoiceResponse, len(question.Choices))
    for i, choice := range question.Choices {
        choices[i] = ToChoiceResponse(choice)
    }

    return PublicQuestionResponse{
        ID:           question.ID,
        QuestionText: question.QuestionText,
        AssessmentID: question.AssessmentID,
        CreatedAt:    question.CreatedAt,
        UpdatedAt:    question.UpdatedAt,
        Choices:      choices,
    }
}

func ToChoiceResponse(choice entities.Choice) PublicChoiceResponse {
    return PublicChoiceResponse{
        ID:         choice.ID,
        ChoiceText: choice.ChoiceText,
        QuestionID: choice.QuestionID,
        CreatedAt:  choice.CreatedAt,
        UpdatedAt:  choice.UpdatedAt,
    }
}

func ToQuestionResponses(questions []entities.Question) []PublicQuestionResponse {
    responses := make([]PublicQuestionResponse, len(questions))
    for i, question := range questions {
        responses[i] = ToQuestionResponse(question)
    }
    return responses
}