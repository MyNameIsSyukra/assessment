package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)
type Answer struct {
	gorm.Model
	Id   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	QuestionId uuid.UUID `gorm:"type:uuid" json:"question_id"`
	ChoiceId   uuid.UUID `gorm:"type:uuid" json:"choice_id"`
	StudentId  uuid.UUID `gorm:"type:uuid" json:"student_id"`
	CreatedAt  time.Time `json:"created_at"`
	Question Question `gorm:"foreignKey:QuestionId" json:"question"`
	Choice Choice `gorm:"foreignKey:ChoiceId" json:"choice"`
}
