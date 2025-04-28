package repository

import (
	entities "assesment/entities"
	"context"

	"gorm.io/gorm"
)

type (
	AssessmentRepository interface {
	CreateAssessment(ctx context.Context, tx *gorm.DB, assesment *entities.Assessment) (*entities.Assessment, error)
	GetAssessmentByID(ctx context.Context, tx *gorm.DB, id string) (*entities.Assessment, error)
	GetAllAssessments() ([]entities.Assessment, error)
	UpdateAssessment(ctx context.Context, tx *gorm.DB,assesment *entities.Assessment) (*entities.Assessment, error)
	DeleteAssessment(ctx context.Context, tx *gorm.DB,id string) error
}
	assesmentRepository struct {
		Db *gorm.DB
	}
)

func NewAssessmentRepository(db *gorm.DB) AssessmentRepository {
	return &assesmentRepository{Db: db}
} 

func (assesmentRepo *assesmentRepository) CreateAssessment(ctx context.Context, tx *gorm.DB,assesment *entities.Assessment) (*entities.Assessment, error) {
	if err := assesmentRepo.Db.Create(assesment).Error; err != nil {
		return nil, err
	}
	return assesment, nil
}

func (assesmentRepo *assesmentRepository) GetAssessmentByID(ctx context.Context, tx *gorm.DB,id string) (*entities.Assessment, error) {
	var assesment entities.Assessment
	if err := assesmentRepo.Db.Where("id = ?", id).First(&assesment).Error; err != nil {
		return nil, err
	}
	return &assesment, nil
}

func (assesmentRepo *assesmentRepository) GetAllAssessments() ([]entities.Assessment, error) {
	var assessments []entities.Assessment
	if err := assesmentRepo.Db.Find(&assessments).Error; err != nil {
		return nil, err
	}

	return assessments, nil
}

func (assesmentRepo *assesmentRepository) UpdateAssessment(ctx context.Context, tx *gorm.DB,assesment *entities.Assessment) (*entities.Assessment, error) {
	if err := assesmentRepo.Db.Where("id = ?", assesment.ID).Where("class_id",assesment.ClassID).Updates(assesment).Error; err != nil {
		return nil, err
	}
	return assesment, nil
}

func (assesmentRepo *assesmentRepository) DeleteAssessment(ctx context.Context, tx *gorm.DB,id string) error {
	if err := assesmentRepo.Db.Delete(&entities.Assessment{},"id",id).Error; err != nil {
		return err
	}
	return nil
}

