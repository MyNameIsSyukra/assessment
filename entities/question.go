package entities

import (
	"github.com/google/uuid"
)

// type Question struct {
// 	Id uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
// 	QuestionText string `json:"question_text"`
// 	EvaluationId string `gorm:"type:uuid" json:"evaluation_id"`
// 	Assessment Assessment `gorm:"foreignKey:EvaluationId;references:Id"`
// }

type Question struct {
    ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
    QuestionText string    `json:"question_text"`
    EvaluationID uuid.UUID `gorm:"type:uuid" json:"evaluation_id"`

    // Relasi ke Assessment
    Assessment Assessment `gorm:"foreignKey:EvaluationID;references:ID" json:"assessment"`

    // Relasi ke Choice (Question has many Choices)
    Choices []Choice `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"choices"`

    // Relasi ke Answer (Question has many Answers)
    Answers []Answer `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"answers"`
}
