package dto

import (
	entities "assesment/entities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnswerCreateRequest struct {
	IdQuestion uuid.UUID `gorm:"type:uuid" binding:"required" json:"question_id"`
	SubmisiionID uuid.UUID `gorm:"type:uuid" binding:"required" json:"submission_id"`
	IdChoice   uuid.UUID `gorm:"type:uuid" binding:"required" json:"choice_id"`
}

type AnswerResponse struct {
	ID         uuid.UUID `gorm:"type:uuid" json:"answer_id"`
	IdQuestion uuid.UUID `gorm:"type:uuid" json:"question_id"`
	SubmissionID uuid.UUID `gorm:"type:uuid" json:"submission_id"`
	IdChoice   uuid.UUID `gorm:"type:uuid" json:"choice_choice"`
	CreatedAt  time.Time `json:"created_at"`
}

type AnswerUpdateRequest struct {
	ID         uuid.UUID `gorm:"type:uuid" json:"answer_id" binding:"required"`
	IdQuestion uuid.UUID `gorm:"type:uuid" json:"question_id"`
	IdStudent  uuid.UUID `gorm:"type:uuid" json:"student_id"`
	IdChoice   uuid.UUID `gorm:"type:uuid" json:"choice_id"`
}

type AnswerUpdatedResponse struct {
	ID         uuid.UUID `gorm:"type:uuid;" json:"answer_id"`
    QuestionID uuid.UUID `gorm:"type:uuid" json:"question_id"`
    ChoiceID   uuid.UUID `gorm:"type:uuid" json:"choice_id"`
    SubmissionID uuid.UUID `gorm:"type:uuid" json:"submission_id"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
    DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type ChoiceGetAnswerBySubmissionIDResponse struct{
    ID         uuid.UUID `gorm:"type:uuid;" json:"choice_id"`
    ChoiceText string    `json:"choice_text"`
    QuestionID uuid.UUID `gorm:"type:uuid" json:"question_id"`
    CreatedAt  time.Time    `json:"created_at"`
    UpdatedAt  time.Time    `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type GetAnswerBySubmissionIDResponse struct {
	ID         uuid.UUID `gorm:"type:uuid" json:"answer_id"`
	SubmissionID uuid.UUID `gorm:"type:uuid" json:"submission_id"`
	IdQuestion uuid.UUID `gorm:"type:uuid" json:"question_id"`
	IdChoice   uuid.UUID `gorm:"type:uuid" json:"choice_id"`
	CreatedAt  time.Time `json:"created_at"`
	Question entities.Question `gorm:"foreignKey:QuestionID;references:ID" json:"question"`
	Choice ChoiceGetAnswerBySubmissionIDResponse `gorm:"foreignKey:ChoiceID;references:ID" json:"choice"`
}
