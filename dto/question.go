package dto

type QuestionCreateRequest struct {
	QuestionText string `json:"question_text"`
	EvaluationId string `gorm:"type:uuid" json:"evaluation_id"`
}

type QuestionUpdateRequest struct {
	Id           string `json:"id"`
	QuestionText string `json:"question_text"`
	EvaluationId string `gorm:"type:uuid" json:"evaluation_id"`
}
