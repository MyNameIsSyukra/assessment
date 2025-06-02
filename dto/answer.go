package dto

import (
	"assesment/entities"
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

type Answer struct {
	ID         uuid.UUID `json:"answer_id"`
	QuestionID uuid.UUID `json:"question_id"`
	ChoiceID   uuid.UUID `json:"choice_id"`
	ChoiceText string `json:"choice_text"`
	SubmissionID uuid.UUID `json:"submission_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type QuestionsSubmittedAnswer struct {
	QuestionID	 uuid.UUID `json:"question_id"`
    QuestionText string    `json:"question_text"`
    AssessmentID uuid.UUID `json:"assessment_id"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
    DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	Choice []entities.Choice `json:"choices"`
	Answers *Answer `json:"answers"`
}

type ContinueSubmissionIDResponse struct {
	Assessment_ID uuid.UUID `json:"assessment_id"`
	SubmissionID uuid.UUID `json:"submission_id"`
	User_ID     uuid.UUID `json:"user_id"`
	EndedTime    time.Time `json:"end_time"`
	Question []QuestionsSubmittedAnswer `json:"questions"`
}
