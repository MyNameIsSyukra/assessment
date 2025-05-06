package dto

import (
	entities "assesment/entities"
	"time"

	"github.com/google/uuid"
)

type AssessmentCreateRequest struct {
	Name string `json:"name" binding:"required"`
	ClassId uuid.UUID `json:"class_id" binding:"required"`
	Duration int `json:"duration" binding:"required"`
	Date_created time.Time `json:"date_created" binding:"required"`
	Start_time time.Time `json:"start_time" binding:"required"`
	End_time time.Time	`json:"end_time" binding:"required"`
}

type AssessmentCreateResponse struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	ClassId uuid.UUID `json:"class_id"`
	Duration int `json:"duration"`
	Date_created time.Time `json:"date_created"`
	Start_time time.Time `json:"start_time"`
	End_time time.Time	`json:"end_time"`
	Updated_At time.Time `json:"updated_at"`
}

type GetAllAssessmentsResponse struct {
	Assessments []entities.Assessment `json:"assessments"`
}

type AssessmentUpdateRequest struct {
	IdEvaluation uuid.UUID `json:"id" binding:"required"`
	Name string `json:"name"`
	ClassId uuid.UUID `json:"class_id"`
	Duration int `json:"duration"`
	Date_created time.Time `json:"date_created"`
	Start_time time.Time `json:"start_time"`
	End_time time.Time	`json:"end_time"`
}


type  StudentGetAllAssesmentByClassIDResponse struct {
	ID        uuid.UUID `gorm:"type:uuid" json:"id"`
    Name      string    `json:"name"`
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
	Duration  int       `json:"duration"`
    CreatedAt time.Time `json:"date_created"`
    UpdatedAt time.Time `json:"updated_at"`
    ClassID   uuid.UUID `gorm:"type:uuid" json:"class_id"`
	SubmissionID *uuid.UUID `json:"submission_id,omitempty"`
	SubmissionStatus entities.ExamStatus  `json:"submission_status"`
}

type GetAssessmentByIDAndByUserIDRequest struct {
	ID uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
}

type GetAssessmentByIDAndByUserIDResponse struct {
	Assessment entities.Assessment `json:"assessment"`
	SubmittedAnswer int `json:"submitted_answer"`
	Question int `json:"question"`
	SubmissionStatus entities.ExamStatus `json:"submission_status"`
	SubmissionID *uuid.UUID `json:"submission_id,omitempty"`
}

// Teacher
type GetAssesmentByIDResponseTeacher struct{
	ID        uuid.UUID `gorm:"type:uuid" json:"id"`
    Name      string    `json:"name"`
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
    Duration int       `json:"duration"` // in Second
	TotalSubmission int `json:"total_submission"`
	TotalStudent int `json:"total_student"`
	Questions []entities.Question `json:"questions"`
}

type GetMemberResponse struct {
	ID uuid.UUID `json:"id"`
	Username string `json:"username"`
	ClassID uuid.UUID `json:"class_id"`
}