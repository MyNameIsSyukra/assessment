package dto

import "time"

type AssessmentCreateRequest struct {
	Name string `json:"name"`
	Date_created time.Time `json:"date_created"`
	Start_time time.Time `json:"start_time"`
	End_time time.Time	`json:"end_time"`
}

type AssessmentUpdateRequest struct {
	IdEvaluation string `json:"id"`
	Name string `json:"name"`
	Date_created time.Time `json:"date_created"`
	Start_time time.Time `json:"start_time"`
	End_time time.Time	`json:"end_time"`
}