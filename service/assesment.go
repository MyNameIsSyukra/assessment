package service

import (
	entities "assesment/entities"
	repository "assesment/repository"
	"context"
)

type (
	AssessmentService interface {
		CreateAssessment(ctx context.Context,assesment *entities.Assessment) (*entities.Assessment, error)
		GetAllAssessments(ctx context.Context)([]entities.Assessment, error)
		GetAssessmentByID(ctx context.Context, id int) (*entities.Assessment, error)
		UpdateAssessment(ctx context.Context, assesment *entities.Assessment) (*entities.Assessment, error)
		DeleteAssessment(ctx context.Context, id int) error
	}
	assesmentService struct {
		assesmentRepo repository.AssessmentRepository
	}
)

func NewAssessmentService(assesmentRepo repository.AssessmentRepository) AssessmentService {
	return &assesmentService{
		assesmentRepo: assesmentRepo,
	}
}

func (assesmentService *assesmentService) CreateAssessment(ctx context.Context, assesment *entities.Assessment) (*entities.Assessment, error) {
	assesment, err := assesmentService.assesmentRepo.CreateAssessment(assesment)
	if err != nil {
		return nil, err
	}
	return assesment, nil
}

func (assesmentService *assesmentService) GetAllAssessments (ctx context.Context) ([]entities.Assessment, error) {
	assessments, err := assesmentService.assesmentRepo.GetAllAssessments()
	if err != nil {
		return nil, err
	}
	return assessments, nil
}

func (assesmentService *assesmentService) GetAssessmentByID (ctx context.Context, id int) (*entities.Assessment,error){
	assesment, err := assesmentService.assesmentRepo.GetAssessmentByID(id)
	if err != nil {
		return nil, err
	}
	return assesment, nil
}

func (assesmentService *assesmentService) UpdateAssessment (ctx context.Context, assesment *entities.Assessment) (*entities.Assessment, error) {
	assesment, err := assesmentService.assesmentRepo.UpdateAssessment(assesment)
	if err != nil {
		return nil, err
	}
	return assesment, nil
}

func (assesmentService *assesmentService) DeleteAssessment (ctx context.Context, id int) error {
	assesment, err := assesmentService.assesmentRepo.GetAssessmentByID(id)
	if err != nil {
		return err
	}

	err = assesmentService.assesmentRepo.DeleteAssessment(int(assesment.ID))
	if err != nil {
		return err
	}
	return nil
}


