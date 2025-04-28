package dto

import (
	entities "assesment/entities"
	"time"

	"github.com/google/uuid"
)

type AssessmentCreateRequest struct {
	Name string `json:"name"`
	ClassId uuid.UUID `json:"class_id"`
	Date_created time.Time `json:"date_created"`
	Start_time time.Time `json:"start_time"`
	End_time time.Time	`json:"end_time"`
}

type AssessmentCreateResponse struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	ClassId uuid.UUID `json:"class_id"`
	Date_created time.Time `json:"date_created"`
	Start_time time.Time `json:"start_time"`
	End_time time.Time	`json:"end_time"`
	Updated_At time.Time `json:"updated_at"`
}

type GetAllAssessmentsResponse struct {
	Assessments []entities.Assessment `json:"assessments"`
}

type AssessmentUpdateRequest struct {
	IdEvaluation string `json:"id"`
	Name string `json:"name"`
	Date_created time.Time `json:"date_created"`
	Start_time time.Time `json:"start_time"`
	End_time time.Time	`json:"end_time"`
}