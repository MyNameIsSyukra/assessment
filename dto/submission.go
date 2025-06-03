package dto

import (
	entities "assesment/entities"
	"time"

	"github.com/google/uuid"
)

type SubmissionCreateRequest struct {
	UserID	   uuid.UUID `json:"user_id" binding:"required"`
	AssessmentID uuid.UUID `json:"assessment_id" binding:"required"`
}

type SubmissionCreateResponse struct {
	ID           uuid.UUID      `gorm:"type:uuid" json:"submission_id"`
	UserID       uuid.UUID      `gorm:"type:uuid" json:"user_id"`
	AssessmentID uuid.UUID      `gorm:"type:uuid" json:"assessment_id"`
	EndedTime	time.Time      `json:"ended_time"`
	Question []PublicQuestionResponse `json:"question"`
}


type GetSubmissionStudentResponse struct {
	ID        *uuid.UUID `gorm:"type:uuid" json:"submission_id"`
	Username  string    `json:"username"`
	User_userID uuid.UUID `gorm:"type:uuid" json:"user_user_id"`
	PhotoUrl string    `json:"photo_url"`
	Role  Role `json:"role"`
	Kelas_kelasID uuid.UUID `gorm:"type:uuid" json:"kelas_kelas_id"`
	Status entities.ExamStatus `json:"status"`
	Score  int       `json:"score"`
	TimeRemaining *int        `json:"time_remaining"`
}

type Role string

const (
	RoleStudent Role = "student"
	RoleTeacher Role = "teacher"
)
