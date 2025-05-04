package repository

import (
	"assesment/dto"
	entities "assesment/entities"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	AssessmentRepository interface {
	CreateAssessment(ctx context.Context, tx *gorm.DB, assesment *entities.Assessment) (*entities.Assessment, error)
	GetAssessmentByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*entities.Assessment, error)
	GetAllAssessments() ([]entities.Assessment, error)
	GetAllAssesmentByClassID(ctx context.Context, tx *gorm.DB,classID uuid.UUID) ([]entities.Assessment, error)
	UpdateAssessment(ctx context.Context, tx *gorm.DB,assesment *entities.Assessment) (*entities.Assessment, error)
	DeleteAssessment(ctx context.Context, tx *gorm.DB,id string) error

	// Student
	StudentGetAllAssesmentByClassIDAndUserID(ctx context.Context, tx *gorm.DB, classID uuid.UUID,userID uuid.UUID) ([]dto.StudentGetAllAssesmentByClassIDResponse, error)
	GetAssessmentByIDAndByUserID(ctx context.Context, tx *gorm.DB, id uuid.UUID, userID uuid.UUID) (*dto.GetAssessmentByIDAndByUserIDResponse, error)
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

func (assesmentRepo *assesmentRepository) GetAssessmentByID(ctx context.Context, tx *gorm.DB,id uuid.UUID) (*entities.Assessment, error) {
	var assesment entities.Assessment
	if err := assesmentRepo.Db.Where("id = ?", id).First(&assesment).Error; err != nil {
		return nil, err
	}
	return &assesment, nil
}

// GetAllAssesmentByClassID retrieves all assessments by class ID
func (assesmentRepo *assesmentRepository) GetAllAssesmentByClassID(ctx context.Context, tx *gorm.DB,classID uuid.UUID) ([]entities.Assessment, error) {
	var assessments []entities.Assessment
	if err := assesmentRepo.Db.Where("class_id = ?", classID).Find(&assessments).Error; err != nil {
		return nil, err
	}

	return assessments, nil
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
// Student
func (assesmentRepo *assesmentRepository) StudentGetAllAssesmentByClassIDAndUserID(ctx context.Context, tx *gorm.DB, classID uuid.UUID,userID uuid.UUID) ([]dto.StudentGetAllAssesmentByClassIDResponse, error) {
	var datas []dto.StudentGetAllAssesmentByClassIDResponse
	var assessments []entities.Assessment
	if err := assesmentRepo.Db.Where("class_id = ?", classID).Find(&assessments).Error; err != nil {
		return nil, err
	}
	for i := range assessments {
		var data dto.StudentGetAllAssesmentByClassIDResponse
		// check if the assessment is submitted by the user
		var submission entities.Submission
		if err := assesmentRepo.Db.Where("assessment_id = ? AND user_id = ?", assessments[i].ID, userID).First(&submission).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return nil, err
			}
		}
		data.ClassID = assessments[i].ClassID
		data.ID = assessments[i].ID
		data.Name = assessments[i].Name
		data.StartTime = assessments[i].StartTime
		data.EndTime = assessments[i].EndTime
		data.CreatedAt = assessments[i].CreatedAt
		data.UpdatedAt = assessments[i].UpdatedAt
		if submission.ID != uuid.Nil {
			data.SubmissionID = &submission.ID
			data.SubmissionStatus = submission.Status
		} else {
			data.SubmissionID = nil
			data.SubmissionStatus = entities.StatusTodo
		}
		datas = append(datas, data)
	}

	return datas, nil
}

func (assesmentRepo *assesmentRepository) GetAssessmentByIDAndByUserID(ctx context.Context, tx *gorm.DB, id uuid.UUID, userID uuid.UUID) (*dto.GetAssessmentByIDAndByUserIDResponse, error) {
	var assessment entities.Assessment
	if err := assesmentRepo.Db.Where("id = ?", id).First(&assessment).Error; err != nil {
		return nil, err
	}
	var submission entities.Submission
	submission.Status = entities.StatusTodo
	if err := assesmentRepo.Db.Where("assessment_id = ? AND user_id = ?", id, userID).First(&submission).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	if submission.ID != uuid.Nil {
		response := &dto.GetAssessmentByIDAndByUserIDResponse{
			Assessment:      assessment,
			SubmissionStatus: submission.Status,
			SubmissionID:    &submission.ID,
		}
		return response, nil
	}
	response := &dto.GetAssessmentByIDAndByUserIDResponse{
		Assessment:      assessment,
		SubmissionStatus: submission.Status,
		SubmissionID:    nil,
	}
	return response, nil
}