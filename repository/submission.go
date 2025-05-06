package repository

import (
	entities "assesment/entities"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	SubmissionRepository interface {
		CreateSubmission(ctx context.Context, tx *gorm.DB, submission *entities.Submission) (*entities.Submission, error)
		GetSubmissionByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*entities.Submission, error)
		GetAllSubmissions() ([]entities.Submission, error)
		UpdateSubmission(ctx context.Context, tx *gorm.DB, submission *entities.Submission) (*entities.Submission, error)
		DeleteSubmission(ctx context.Context, tx *gorm.DB, id uuid.UUID) error
		GetSubmissionsByAssessmentID(ctx context.Context, tx *gorm.DB, assessmentID uuid.UUID) ([]entities.Submission, error)
		GetSubmissionsByUserID(ctx context.Context, tx *gorm.DB, userID uuid.UUID) ([]entities.Submission, error)
		GetSubmissionsByAssessmentIDAndUserID(ctx context.Context, tx *gorm.DB, assessmentID uuid.UUID, userID uuid.UUID) (*entities.Submission, bool,error)
		Submitted(ctx context.Context, tx *gorm.DB, submission *entities.Submission) (*entities.Submission, error)
	}
	submissionRepository struct {
		Db *gorm.DB
	}
)


func NewSubmissionRepository(db *gorm.DB) SubmissionRepository {
	return &submissionRepository{Db: db}
}


func (submissionRepo *submissionRepository) CreateSubmission(ctx context.Context, tx *gorm.DB, submission *entities.Submission) (*entities.Submission, error) {
	submissionExist := submissionRepo.Db.Where("id = ?", submission.ID).First(&submission)
	if submissionExist.Error == nil {
		return nil, errors.New("submission already exists")
	}
	
	if err := submissionRepo.Db.Create(submission).Error; err != nil {
		return nil, err
	}
	return submission, nil
}

func (submissionRepo *submissionRepository) GetSubmissionByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*entities.Submission, error) {
	var submission entities.Submission
	if err := submissionRepo.Db.Where("id = ?", id).First(&submission).Error; err != nil {
		return nil, err
	}
	return &submission, nil
}

func (submissionRepo *submissionRepository) GetAllSubmissions() ([]entities.Submission, error) {
	var submissions []entities.Submission
	if err := submissionRepo.Db.Find(&submissions).Error; err != nil {
		return nil, err
	}
	return submissions, nil
}

func (submissionRepo *submissionRepository) UpdateSubmission(ctx context.Context, tx *gorm.DB, submission *entities.Submission) (*entities.Submission, error) {
	
	
	if err := submissionRepo.Db.Save(submission).Error; err != nil {
		return nil, err
	}
	return submission, nil
}		

func (submissionRepo *submissionRepository) DeleteSubmission(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	var submission entities.Submission
	if err := submissionRepo.Db.Where("id = ?", id).First(&submission).Error; err != nil {
		return err
	}
	if err := submissionRepo.Db.Delete(&submission).Error; err != nil {
		return err
	}
	return nil
}

func (submissionRepo *submissionRepository) GetSubmissionsByAssessmentID(ctx context.Context, tx *gorm.DB, assessmentID uuid.UUID) ([]entities.Submission, error) {
	var submissions []entities.Submission
	if err := submissionRepo.Db.Where("assessment_id = ?", assessmentID).Preload("Assessment").Find(&submissions).Error; err != nil {
		return nil, err
	}
	return submissions, nil
}		

func (submissionRepo *submissionRepository) GetSubmissionsByUserID(ctx context.Context, tx *gorm.DB, userID uuid.UUID) ([]entities.Submission, error) {
	var submissions []entities.Submission
	if err := submissionRepo.Db.Where("user_id = ?", userID).Find(&submissions).Error; err != nil {
		return nil, err
	}
	return submissions, nil
}

func (submissionRepo *submissionRepository) GetSubmissionsByAssessmentIDAndUserID(ctx context.Context, tx *gorm.DB, assessmentID uuid.UUID, userID uuid.UUID) (*entities.Submission,bool, error) {
	var submission entities.Submission
	if err := submissionRepo.Db.Where("assessment_id = ? AND user_id = ?", assessmentID, userID).First(&submission).Error; err != nil {
		return nil, false, err
	}
	return &submission, true, nil
}

func (submissionRepo *submissionRepository) Submitted(ctx context.Context, tx *gorm.DB, submission *entities.Submission) (*entities.Submission, error) {
	var countIsCorrect int64
	var totalQuestions int64
	existingSubmission := submissionRepo.Db.Where("id = ?", submission.ID)
	if existingSubmission.Error != nil {
		return nil, errors.New("submission not found")
	}

	if tx == nil {
		tx = submissionRepo.Db
	}
	
	err := tx.Model(&entities.Answer{}).
		Joins("JOIN choices ON answers.choice_id = choices.id").
		Where("answers.submission_id = ? AND choices.is_correct = ?", submission.ID, true).
		Count(&countIsCorrect).Error
	if err != nil {
		return nil, err
	}
	err = tx.Model(&entities.Question{}).
	    Where("evaluation_id = ?", submission.AssessmentID).
	    Count(&totalQuestions).Error
	if err != nil {
	    return nil, err
	}

	print("count", countIsCorrect)
	print("total", totalQuestions)
	submission.Score = float64(countIsCorrect) / float64(totalQuestions) * 100
	// submission.Score = score
	if err := submissionRepo.Db.Save(submission).Error; err != nil {
		return nil, err
	}


	return submission, nil
}


