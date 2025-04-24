package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Choice struct {
	gorm.Model
	Id     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	ChoiceText   string    `json:"choice_text"`
	QuestionId   uuid.UUID `gorm:"type:uuid" json:"question_id"`
	IsCorrect    bool      `json:"is_correct"`
	Question     Question  `gorm:"foreignKey:QuestionId;references:IdQuestion"`
}