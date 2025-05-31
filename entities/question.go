package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Question struct {
    ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"question_id"`
    QuestionText string    `json:"question_text"`
    AssessmentID uuid.UUID `gorm:"type:uuid" json:"assessment_id"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
    DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
    // Relasi ke Assessment
    Assessment Assessment `gorm:"foreignKey:AssessmentID;references:ID" json:"-"`

    // Relasi ke Choice (Question has many Choices)
    Choices []Choice `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"choice"`

    // Relasi ke Answer (Question has many Answers)
    Answers []Answer `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}
