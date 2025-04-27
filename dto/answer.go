package dto

import "github.com/google/uuid"

type AnswerCreateRequest struct {
	IdQuestion uuid.UUID `gorm:"type:uuid" json:"id_question"`
	IdStudent  uuid.UUID `gorm:"type:uuid" json:"id_student"`
	IdChoice   uuid.UUID `gorm:"type:uuid" json:"id_choice"`
}

type AnswerUpdateRequest struct {
	ID         string ` json:"id"`
	IdQuestion uuid.UUID `gorm:"type:uuid" json:"id_question"`
	IdStudent  uuid.UUID `gorm:"type:uuid" json:"id_student"`
	IdChoice   uuid.UUID `gorm:"type:uuid" json:"id_choice"`
}