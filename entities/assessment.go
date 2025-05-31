package entities

import (
	"time"

	"github.com/google/uuid"
)

// type Assessment struct{
// 	Id uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
// 	Name string `json:"name"`
// 	CreatedAt time.Time `gorm:"colomn:created_at" json:"date_created"`
// 	Start_time time.Time `json:"start_time"`
// 	End_time time.Time `json:"end_time"`
// 	UpdatedAt time.Time `gorm:"colomn:updated_at" json:"updated_at`
// 	ClassId uuid.UUID `gorm:"type:uuid" json:"class_id"`
// }

type Assessment struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"assessment_id"`
    Name      string    `json:"name"`
    Description string `json:"description"`
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
    Duration int       `json:"duration"` // in Second
    CreatedAt time.Time `json:"date_created"`
    UpdatedAt time.Time `json:"updated_at"`
    ClassID   uuid.UUID `gorm:"type:uuid" json:"class_id"`

    // Relasi ke Question (Assessment has many Questions)
    Questions []Question `gorm:"foreignKey:AssessmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"questions,omitempty"`
    // Relasi ke Submission (Assessment has many Submissions)
    Submissions []Submission `gorm:"foreignKey:AssessmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"submissions,omitempty"`
}
