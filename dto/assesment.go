package dto

import (
	entities "assesment/entities"
	"time"

	"github.com/google/uuid"
)

type AssessmentCreateRequest struct {
	Name string `json:"name" binding:"required"`
	ClassId uuid.UUID `json:"class_id" binding:"required"`
	Date_created time.Time `json:"date_created" binding:"required"`
	Start_time time.Time `json:"start_time" binding:"required"`
	End_time time.Time	`json:"end_time" binding:"required"`
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
	IdEvaluation uuid.UUID `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Date_created time.Time `json:"date_created" binding:"required"`
	Start_time time.Time `json:"start_time" binding:"required"`
	End_time time.Time	`json:"end_time" binding:"required"`
}