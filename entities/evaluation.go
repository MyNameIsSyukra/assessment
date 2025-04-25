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
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"date_created"`
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
    UpdatedAt time.Time `json:"updated_at"`
    ClassID   uuid.UUID `gorm:"type:uuid" json:"class_id"`
}
