package repository

import (
	entities "assesment/entities"

	"gorm.io/gorm"
)

type (
	AssessmentRepository interface {
	CreateAssessment(assesment *entities.Assessment) (*entities.Assessment, error)
	GetAssessmentByID(id int) (*entities.Assessment, error)
	GetAllAssessments() ([]entities.Assessment, error)
	UpdateAssessment(assesment *entities.Assessment) (*entities.Assessment, error)
	DeleteAssessment(id int) error
}
	assesmentRepository struct {
		Db *gorm.DB
	}
)

func NewAssessmentRepository(db *gorm.DB) AssessmentRepository {
	return &assesmentRepository{Db: db}
} 

func (assesmentRepo *assesmentRepository) CreateAssessment(assesment *entities.Assessment) (*entities.Assessment, error) {
	if err := assesmentRepo.Db.Create(assesment).Error; err != nil {
		return nil, err
	}
	return assesment, nil
}

func (assesmentRepo *assesmentRepository) GetAssessmentByID(id int) (*entities.Assessment, error) {
	var assesment entities.Assessment
	if err := assesmentRepo.Db.First(&assesment, id).Error; err != nil {
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

func (assesmentRepo *assesmentRepository) UpdateAssessment(assesment *entities.Assessment) (*entities.Assessment, error) {
	if err := assesmentRepo.Db.Save(assesment).Error; err != nil {
		return nil, err
	}
	return assesment, nil
}

func (assesmentRepo *assesmentRepository) DeleteAssessment(id int) error {
	var assesment entities.Assessment
	if err := assesmentRepo.Db.Delete(&assesment, id).Error; err != nil {
		return err
	}
	return nil
}

