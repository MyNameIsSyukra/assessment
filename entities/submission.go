package entities

import (
	"time"

	"github.com/google/uuid"
)

type Submission struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID       uuid.UUID      `gorm:"type:uuid" json:"user_id"`
	AssessmentID uuid.UUID      `gorm:"type:uuid" json:"assessment_id"`
	EndedTime	time.Time      `json:"ended_time"`
	SubmittedAt    time.Time      `json:"submitted_at"`
	Score	   float64       `json:"score"`
	Status	   ExamStatus       `json:"status"`

	Assessment Assessment `gorm:"foreignKey:AssessmentID;references:ID" json:"assessment,omitempty"`
	Answers    []Answer    `gorm:"foreignKey:SubmissionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"answers,omitempty"`
}

type ExamStatus string


const (
	StatusInProgress ExamStatus = "in_progress"
	StatusSubmitted  ExamStatus = "submitted"
	StatusTodo 	 ExamStatus = "todo"
)