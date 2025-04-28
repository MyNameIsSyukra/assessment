package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// type Answer struct {
// 	Id   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
// 	QuestionId uuid.UUID `gorm:"type:uuid" json:"question_id"`
// 	ChoiceId   uuid.UUID `gorm:"type:uuid" json:"choice_id"`
// 	StudentId  uuid.UUID `gorm:"type:uuid" json:"student_id"`
// 	CreatedAt  time.Time `json:"created_at"`
// 	Question Question `gorm:"foreignKey:QuestionId;references:Id" json:"question"`
// 	Choice Choice `gorm:"foreignKey:ChoiceId;references:Id" json:"choice"`
// }

type Answer struct {
    ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
    QuestionID uuid.UUID `gorm:"type:uuid" json:"question_id"`
    ChoiceID   uuid.UUID `gorm:"type:uuid" json:"choice_id"`
    StudentID  uuid.UUID `gorm:"type:uuid" json:"student_id"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
    DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
    // Relasi ke Question
    Question Question `gorm:"foreignKey:QuestionID;references:ID" json:"question"`

    // Relasi ke Choice
    Choice Choice `gorm:"foreignKey:ChoiceID;references:ID" json:"choice"`
}
