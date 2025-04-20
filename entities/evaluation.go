package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Evaluation struct{
	gorm.Model
	IdEvaluation uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	name string `json:"name"`
	date_created time.Time `json:"date_created"`
	start_time time.Time `json:"start_time"`
	end_time time.Time `json:"end_time"`
}