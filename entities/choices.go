package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// type Choice struct {
// 	gorm.Model
// 	Id     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
// 	ChoiceText   string    `json:"choice_text"`
// 	QuestionId   uuid.UUID `gorm:"type:uuid" json:"question_id"`
// 	IsCorrect    bool      `json:"is_correct"`
// 	Question     Question  `gorm:"foreignKey:QuestionId;references:Id"`
// }

type Choice struct {
    ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
    ChoiceText string    `json:"choice_text"`
    QuestionID uuid.UUID `gorm:"type:uuid" json:"question_id"`
    IsCorrect  bool      `json:"is_correct"`
    CreatedAt  time.Time    `json:"created_at"`
    UpdatedAt  time.Time    `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`


    // Relasi ke Question
    Question Question `gorm:"foreignKey:QuestionID;references:ID" json:"-"`

    // Relasi ke Answer (Choice has many Answers)
    Answers []Answer `gorm:"foreignKey:ChoiceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}
