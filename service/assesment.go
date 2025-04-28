package service

import (
	dto "assesment/dto"
	entities "assesment/entities"
	repository "assesment/repository"
	"context"
)

type (
	AssessmentService interface {
		CreateAssessment(ctx context.Context,assesment *dto.AssessmentCreateRequest) (dto.AssessmentCreateResponse, error)
		GetAllAssessments(ctx context.Context)(dto.GetAllAssessmentsResponse, error)
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

func (assesmentService *assesmentService) CreateAssessment(ctx context.Context, assesment *dto.AssessmentCreateRequest) (dto.AssessmentCreateResponse, error) {
	assesmentEntity := entities.Assessment{
		ClassID: assesment.ClassId,
		Name: assesment.Name,
		CreatedAt: assesment.Date_created,
		StartTime: assesment.Start_time,
		EndTime: assesment.End_time,
	}
	
	createdAssessment, err := assesmentService.assesmentRepo.CreateAssessment(ctx, nil, &assesmentEntity)
	if err != nil {
		return dto.AssessmentCreateResponse{}, err
	}
	res := dto.AssessmentCreateResponse{
		ID: createdAssessment.ID,
		Name: createdAssessment.Name,
		ClassId: createdAssessment.ClassID,
		Start_time: createdAssessment.StartTime,
		End_time: createdAssessment.EndTime,
		Date_created: createdAssessment.CreatedAt,
		Updated_At: createdAssessment.UpdatedAt,
	}
	return res, nil	
}

func (assesmentService *assesmentService) GetAllAssessments (ctx context.Context) (dto.GetAllAssessmentsResponse, error) {
	assessments, err := assesmentService.assesmentRepo.GetAllAssessments()
	if err != nil {
		return dto.GetAllAssessmentsResponse{}, err
	}
	return dto.GetAllAssessmentsResponse{
		Assessments: assessments,
	}, nil
}

func (assesmentService *assesmentService) GetAssessmentByID (ctx context.Context, id string) (*entities.Assessment,error){
	assesment, err := assesmentService.assesmentRepo.GetAssessmentByID(ctx,nil,id)
	if err != nil {
		return nil, err
	}
	return assesment, nil
}

func (assesmentService *assesmentService) UpdateAssessment (ctx context.Context, assesment *dto.AssessmentUpdateRequest) (*entities.Assessment, error) {
	ass,err := assesmentService.assesmentRepo.GetAssessmentByID(ctx, nil, assesment.IdEvaluation)
	print(ass)
	if ass == nil {
		return nil, err
	}
	assesmentEntity := entities.Assessment{
		ID: ass.ID,
		ClassID: ass.ClassID,
		Name: assesment.Name,
		CreatedAt: assesment.Date_created,
		StartTime: assesment.Start_time,
		EndTime: assesment.End_time,
	}
	
	updatedAssessment, err := assesmentService.assesmentRepo.UpdateAssessment(ctx, nil, &assesmentEntity)
	if err != nil {
		return nil, err
	}
	return updatedAssessment, nil
}

func (assesmentService *assesmentService) DeleteAssessment (ctx context.Context, id string) error {
	asses,err := assesmentService.assesmentRepo.GetAssessmentByID(ctx, nil, id)
	if err != nil {
		return err
	}

	err = assesmentService.assesmentRepo.DeleteAssessment(ctx, nil,asses.ID.String())
	if err != nil {
		return err
	}
	return nil
}


