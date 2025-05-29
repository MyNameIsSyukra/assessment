package dto

import (
	entities "assesment/entities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnswerCreateRequest struct {
	IdQuestion uuid.UUID `gorm:"type:uuid" binding:"required" json:"id_question"`
	SubmisiionID uuid.UUID `gorm:"type:uuid" binding:"required" json:"submission_id"`
	IdChoice   uuid.UUID `gorm:"type:uuid" binding:"required" json:"id_choice"`
}

type AnswerResponse struct {
	ID         uuid.UUID `gorm:"type:uuid" json:"id"`
	IdQuestion uuid.UUID `gorm:"type:uuid" json:"id_question"`
	// IdStudent  uuid.UUID `gorm:"type:uuid" json:"id_student"`
	SubmissionID uuid.UUID `gorm:"type:uuid" json:"submission_id"`
	IdChoice   uuid.UUID `gorm:"type:uuid" json:"id_choice"`
	CreatedAt  time.Time `json:"created_at"`
}

type AnswerUpdateRequest struct {
	ID         uuid.UUID `gorm:"type:uuid" json:"id" binding:"required"`
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

type ChoiceGetAnswerBySubmissionIDResponse struct{
    ID         uuid.UUID `gorm:"type:uuid;" json:"choice_id"`
    ChoiceText string    `json:"choice_text"`
    QuestionID uuid.UUID `gorm:"type:uuid" json:"question_id"`
    CreatedAt  time.Time    `json:"created_at"`
    UpdatedAt  time.Time    `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
type GetAnswerBySubmissionIDResponse struct {
	ID         uuid.UUID `gorm:"type:uuid" json:"id_answer"`
	SubmissionID uuid.UUID `gorm:"type:uuid" json:"submission_id"`
	IdQuestion uuid.UUID `gorm:"type:uuid" json:"id_question"`
	IdChoice   uuid.UUID `gorm:"type:uuid" json:"id_choice"`
	CreatedAt  time.Time `json:"created_at"`
	Question entities.Question `gorm:"foreignKey:QuestionID;references:ID" json:"question"`
	Choice ChoiceGetAnswerBySubmissionIDResponse `gorm:"foreignKey:ChoiceID;references:ID" json:"choice"`
}
