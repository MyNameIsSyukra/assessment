package service

import (
	dto "assesment/dto"
	entities "assesment/entities"
	repository "assesment/repository"
	"context"
)

type (
	AssessmentService interface {
		CreateAssessment(ctx context.Context,assesment *dto.AssessmentCreateRequest) (*entities.Assessment, error)
		GetAllAssessments(ctx context.Context)([]entities.Assessment, error)
		GetAssessmentByID(ctx context.Context, id string) (*entities.Assessment, error)
		UpdateAssessment(ctx context.Context, assesment *dto.AssessmentUpdateRequest) (*entities.Assessment, error)
		DeleteAssessment(ctx context.Context, id string) error
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

func (assesmentService *assesmentService) CreateAssessment(ctx context.Context, assesment *dto.AssessmentCreateRequest) (*entities.Assessment, error) {
	assesmentEntity := entities.Assessment{
		Name: assesment.Name,
		CreatedAt: assesment.Date_created,
		Start_time: assesment.Start_time,
		End_time: assesment.End_time,
	}
	
	createdAssessment, err := assesmentService.assesmentRepo.CreateAssessment(ctx, nil, &assesmentEntity)
	if err != nil {
		return nil, err
	}
	return createdAssessment, nil	
}

func (assesmentService *assesmentService) GetAllAssessments (ctx context.Context) ([]entities.Assessment, error) {
	assessments, err := assesmentService.assesmentRepo.GetAllAssessments()
	if err != nil {
		return nil, err
	}
	return assessments, nil
}

func (assesmentService *assesmentService) GetAssessmentByID (ctx context.Context, id string) (*entities.Assessment,error){
	assesment, err := assesmentService.assesmentRepo.GetAssessmentByID(ctx,nil,id)
	if err != nil {
		return nil, err
	}
	return assesment, nil
}

func (assesmentService *assesmentService) UpdateAssessment (ctx context.Context, assesment *dto.AssessmentUpdateRequest) (*entities.Assessment, error) {
	assesmentEntity := entities.Assessment{
		Name: assesment.Name,
		CreatedAt: assesment.Date_created,
		Start_time: assesment.Start_time,
		End_time: assesment.End_time,
	}
	
	updatedAssessment, err := assesmentService.assesmentRepo.UpdateAssessment(ctx, nil, &assesmentEntity)
	if err != nil {
		return nil, err
	}
	return updatedAssessment, nil
}

func (assesmentService *assesmentService) DeleteAssessment (ctx context.Context, id string) error {
	err := assesmentService.assesmentRepo.DeleteAssessment(ctx, nil, id)
	if err != nil {
		return err
	}
	return nil
}


