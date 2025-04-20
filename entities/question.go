package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Question struct {
	gorm.Model
	IdQuestion uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	QuestionText string `json:"question_text"`
}