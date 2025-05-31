package repository

import (
	"assesment/dto"
	entities "assesment/entities"
	"context"
	"math"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	AssessmentRepository interface {
	CreateAssessment(ctx context.Context, tx *gorm.DB, assesment *entities.Assessment) (*entities.Assessment, error)
	GetAssessmentByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*entities.Assessment, error)
	// GetAllAssessments() ([]entities.Assessment, error)
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
	if err := assesmentRepo.Db.Where("id = ?", id).Preload("Questions").First(&assesment).Error; err != nil {
		return &entities.Assessment{}, err
	}
	return &assesment, nil
}

// GetAllAssesmentByClassID retrieves all assessments by class ID
func (assesmentRepo *assesmentRepository) GetAllAssesmentByClassID(ctx context.Context, tx *gorm.DB,classID uuid.UUID) ([]entities.Assessment, error) {
	var assessments []entities.Assessment
	if err := assesmentRepo.Db.Where("class_id = ?", classID).Find(&assessments).Error; err != nil {
		return []entities.Assessment{}, err
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
		return[]dto.StudentGetAllAssesmentByClassIDResponse{}, err
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
		data.Duration = assessments[i].Duration
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
	if err := assesmentRepo.Db.Where("id = ?", id).Preload("Questions").First(&assessment).Error; err != nil {
		return &dto.GetAssessmentByIDAndByUserIDResponse{}, err
	}
	var submission entities.Submission
	submission.Status = entities.StatusTodo
	if err := assesmentRepo.Db.Where("assessment_id = ? AND user_id = ?", id, userID).Preload("Answers").First(&submission).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return &dto.GetAssessmentByIDAndByUserIDResponse{}, err
		}
	}
	assessmentWithoutQuestions := entities.Assessment{
		ID:          assessment.ID,
		Name:        assessment.Name,
		Description: assessment.Description,
		StartTime:   assessment.StartTime,
		EndTime:     assessment.EndTime,
		Duration:    assessment.Duration,
		CreatedAt:   assessment.CreatedAt,
		UpdatedAt:   assessment.UpdatedAt,
		ClassID:     assessment.ClassID,
	}
	var response dto.GetAssessmentByIDAndByUserIDResponse
	if submission.ID == uuid.Nil {
		response.Assessment = assessmentWithoutQuestions
		response.MaxScore = 100
		response.Question = len(assessment.Questions)
		response.SubmittedAnswer = len(submission.Answers)
		response.Score = nil
		response.SubmissionID = nil
		response.SubmissionStatus = submission.Status
		response.TimeRemaining = nil
		response.TimeSpent = nil
	}else if submission.Status == entities.StatusInProgress{
		timeremain := submission.EndedTime.Sub(time.Now())
		timeremain = timeremain.Round(time.Second)
		if timeremain < 0 {
			timeremain = 0
		}
		timeSpent := time.Now().Sub(submission.CreatedAt)
		timeSpent = timeSpent.Round(time.Second)
		response.Assessment = assessmentWithoutQuestions
		response.MaxScore = 100
		response.Question = len(assessment.Questions)
		response.SubmittedAnswer = len(submission.Answers)
		response.Score = nil
		response.SubmissionID = &submission.ID
		response.SubmissionStatus = submission.Status
		response.TimeRemaining = &timeremain
		response.TimeSpent = &timeSpent
	} else if submission.Status == entities.StatusSubmitted {
		score := int(math.Round(submission.Score))
		timeremain := submission.EndedTime.Sub(submission.CreatedAt)
		timeremain = timeremain.Round(time.Second)
		if timeremain < 0 {
			timeremain = 0
		}
		timeSpent := submission.UpdatedAt.Sub(submission.CreatedAt)
		timeSpent = timeSpent.Round(time.Second)
		response.Assessment = assessmentWithoutQuestions
		response.MaxScore = 100
		response.Question = len(assessment.Questions)
		response.SubmittedAnswer = len(submission.Answers)
		response.Score = &score
		response.SubmissionID = &submission.ID
		response.SubmissionStatus = submission.Status
		response.TimeRemaining = &timeremain
		response.TimeSpent = &timeSpent
	}
	return &response, nil
}
