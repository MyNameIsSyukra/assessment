package dto

import (
	entities "assesment/entities"
	"time"

	"github.com/google/uuid"
)

type AnswerCreateRequest struct {
	IdQuestion uuid.UUID `gorm:"type:uuid" binding:"required" json:"id_question"`
	SubmisiionID uuid.UUID `gorm:"type:uuid" binding:"required" json:"submission_id"`
	IdChoice   uuid.UUID `gorm:"type:uuid" binding:"required" json:"id_choice"`
}

type AnswerResponse struct {
	ID         uuid.UUID `gorm:"type:uuid" json:"id"`
	IdQuestion uuid.UUID `gorm:"type:uuid" json:"id_question"`
	IdStudent  uuid.UUID `gorm:"type:uuid" json:"id_student"`
	SubmissionID uuid.UUID `gorm:"type:uuid" json:"submission_id"`
	IdChoice   uuid.UUID `gorm:"type:uuid" json:"id_choice"`
	CreatedAt  time.Time `json:"created_at"`
}

type AnswerUpdateRequest struct {
	ID         uuid.UUID ` json:"id"`
	IdQuestion uuid.UUID `gorm:"type:uuid" json:"id_question"`
	IdStudent  uuid.UUID `gorm:"type:uuid" json:"id_student"`
	IdChoice   uuid.UUID `gorm:"type:uuid" json:"id_choice"`
}

type GetAnswerByStudentIDRequest struct {
	IdStudent string `json:"id_student" binding:"required"`
	IdAssesment string `json:"id_assesment" binding:"required"`
}

type GetAnswerByStudentIDResponse struct {
	ID         uuid.UUID `gorm:"type:uuid" json:"id"`
    QuestionID uuid.UUID `gorm:"type:uuid" json:"question_id"`
    ChoiceID   uuid.UUID `gorm:"type:uuid" json:"choice_id"`
    StudentID  uuid.UUID `gorm:"type:uuid" json:"student_id"`
    CreatedAt  time.Time `json:"created_at"`

    // Relasi ke Question
    Question entities.Question `gorm:"foreignKey:QuestionID;references:ID" json:"question"`
}

