package service

import (
	dto "assesment/dto"
	entities "assesment/entities"
	repository "assesment/repository"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type (
	SubmissionService interface {
		// student
		CreateSubmission(ctx context.Context, submission *dto.SubmissionCreateRequest) (dto.SubmissionCreateResponse, error)
		Submitted(ctx context.Context, submissionID uuid.UUID) (*entities.Submission, error)
		GetAllSubmissions(ctx context.Context) ([]entities.Submission, error)
		GetSubmissionByID(ctx context.Context, id uuid.UUID) (*entities.Submission, error)
		// UpdateSubmission(ctx context.Context, submission *dto.SubmissionCreateRequest) (*entities.Submission, error)
		DeleteSubmission(ctx context.Context, id uuid.UUID) error
		GetSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]entities.Submission, error)
		GetSubmissionsByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Submission, error)
		GetSubmissionsByAssessmentIDAndUserID(ctx context.Context, assessmentID uuid.UUID, userID uuid.UUID) (*entities.Submission, error)
		GetStudentSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID, flag string)  ([]dto.GetSubmissionStudentResponse, error)
	}
	submissionService struct {
		submissionRepo repository.SubmissionRepository
		questionRepo  repository.QuestionRepository
		assessmentRepo repository.AssessmentRepository
	}
)

func NewSubmissionService(submissionRepo repository.SubmissionRepository, questionRepo repository.QuestionRepository, assessmentRepo repository.AssessmentRepository) SubmissionService {
	return &submissionService{
		submissionRepo: submissionRepo,
		questionRepo:  questionRepo,
		assessmentRepo: assessmentRepo,
	}
}

func (submissionService *submissionService) CreateSubmission(ctx context.Context, submission *dto.SubmissionCreateRequest) (dto.SubmissionCreateResponse, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	assesment, errr := submissionService.assessmentRepo.GetAssessmentByID(ctx, nil, submission.AssessmentID)
	if errr != nil {
		return dto.SubmissionCreateResponse{}, errr
	}

	// check if user is a member of the class
	params := url.Values{}
	params.Add("classID", assesment.ClassID.String())
	params.Add("userID", submission.UserID.String())
	url := fmt.Sprintf("%s/service/class/member/?%s", os.Getenv("CLASS_SERVICE_URL"), params.Encode())
	resp, err := http.Get(url)
	if err != nil {
		return dto.SubmissionCreateResponse{}, fmt.Errorf("failed to get class member: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return dto.SubmissionCreateResponse{}, fmt.Errorf("you are not a member of this class")
	}
	defer resp.Body.Close()

	// check if the submission already exists
	_,flag, _ := submissionService.submissionRepo.GetSubmissionsByAssessmentIDAndUserID(ctx, nil, submission.AssessmentID,submission.UserID)
	if flag {
		return dto.SubmissionCreateResponse{}, errors.New("submission already exists")
	}
	submissionEntity := entities.Submission{
		UserID: 	 submission.UserID,
		AssessmentID: submission.AssessmentID,
		Status: "in_progress",
		EndedTime: time.Now().Add(time.Duration(assesment.Duration) * time.Second),
	}
	
	var question []entities.Question
	question, err = submissionService.questionRepo.GetQuestionsByAssessmentID(ctx, nil, submission.AssessmentID)
	if err != nil {
		question = nil
	}

	createdSubmission, err := submissionService.submissionRepo.CreateSubmission(ctx, nil, &submissionEntity)
	if err != nil {
		return dto.SubmissionCreateResponse{}, err
	}

	return dto.SubmissionCreateResponse{
		ID:             createdSubmission.ID,
		UserID:         createdSubmission.UserID,
		AssessmentID:   createdSubmission.AssessmentID,
		EndedTime:      createdSubmission.EndedTime,
		Question: question,
	}, nil
}

func (submissionService *submissionService) GetAllSubmissions(ctx context.Context) ([]entities.Submission, error) {
	if submissions, err := submissionService.submissionRepo.GetAllSubmissions(); err != nil {
		return []entities.Submission{}, err
	} else {
		return submissions, nil
	}
}

func (submissionService *submissionService) GetSubmissionByID(ctx context.Context, id uuid.UUID) (*entities.Submission, error) {
	submission, err := submissionService.submissionRepo.GetSubmissionByID(ctx, nil, id)
	if err != nil {
		return &entities.Submission{}, err
	}
	return submission, nil
}

func (submissionService *submissionService) DeleteSubmission(ctx context.Context, id uuid.UUID) error {
	if err := submissionService.submissionRepo.DeleteSubmission(ctx, nil, id); err != nil {
		return err
	}
	return nil
}

func (submissionService *submissionService) GetSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]entities.Submission, error) {
	submissions, err := submissionService.submissionRepo.GetSubmissionsByAssessmentID(ctx, nil, assessmentID)
	if err != nil {
		return []entities.Submission{}, err
	}
	return submissions, nil
}

func (submissionService *submissionService) GetSubmissionsByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Submission, error) {
	submissions, err := submissionService.submissionRepo.GetSubmissionsByUserID(ctx, nil, userID)
	if err != nil {
		return []entities.Submission{}, err
	}
	return submissions, nil
}

func (submissionService *submissionService) GetSubmissionsByAssessmentIDAndUserID(ctx context.Context, assessmentID uuid.UUID, userID uuid.UUID) (*entities.Submission, error) {
	submission,_, err := submissionService.submissionRepo.GetSubmissionsByAssessmentIDAndUserID(ctx, nil, assessmentID, userID)
	if err != nil {
		return &entities.Submission{}, err
	}
	return submission, nil
}

func (submissionService *submissionService) Submitted(ctx context.Context, submissionID uuid.UUID) (*entities.Submission, error) {
	submission, err := submissionService.submissionRepo.GetSubmissionByID(ctx, nil, submissionID)
	if err != nil {
		return &entities.Submission{}, err
	}
	data := entities.Submission{
		ID: submission.ID,
		UserID: submission.UserID,
		AssessmentID: submission.AssessmentID,
		Status: "submitted",
	}
	result, err := submissionService.submissionRepo.Submitted(ctx, nil, &data)
	if err != nil {
		return &entities.Submission{}, err
	}

	return result, nil
}

func (submissionService *submissionService) GetStudentSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID,flag string) ([]dto.GetSubmissionStudentResponse, error) {
// func (submissionService *submissionService) GetStudentSubmissionsByAssessmentID(ctx context.Context, assessmentID uuid.UUID,flag string) (, error) {
	assesment, err := submissionService.assessmentRepo.GetAssessmentByID(ctx, nil, assessmentID)
	if err != nil {
		return []dto.GetSubmissionStudentResponse{}, err
	}
	submissions, err := submissionService.submissionRepo.GetSubmissionsByAssessmentID(ctx, nil, assessmentID)
	if err != nil {
		return []dto.GetSubmissionStudentResponse{}, err
	}
	classID := assesment.ClassID
	url := fmt.Sprintf("http://localhost:8081/service/class/%s", classID)
	resp, err := http.Get(url)
	if err != nil {
		return []dto.GetSubmissionStudentResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []dto.GetSubmissionStudentResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}
	var members []dto.GetSubmissionStudentResponse
	if err := json.Unmarshal(body, &members); err != nil {
		return []dto.GetSubmissionStudentResponse{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	
	submissionMap := make(map[uuid.UUID]entities.Submission)
	for _, s := range submissions {
		// Gunakan user_id sebagai key (kalau ada banyak submission per user, bisa dihandle pakai slice)
		submissionMap[s.UserID] = s
	}
	
	status := entities.ExamStatus("todo")
	var studentSubmission []dto.GetSubmissionStudentResponse
	for _, m := range members {
		var datas dto.GetSubmissionStudentResponse
		if m.Role == dto.RoleTeacher{
			continue
		}
		sub := submissionMap[m.User_userID]
		now := time.Now()
		duration := int64(sub.EndedTime.Sub(now).Seconds())
		if duration < 0 {
			duration = 0
		}
		if sub.Status == "" {
			datas.ID = nil
			datas.Username = m.Username
			datas.User_userID = m.User_userID
			datas.Kelas_kelasID = m.Kelas_kelasID
			datas.Role = m.Role
			datas.Status = status	
			datas.TimeRemaining = nil
			datas.Score = 0
		}else if sub.Status == "in_progress" {
			datas.ID = &sub.ID
			datas.Username = m.Username
			datas.User_userID = m.User_userID
			datas.Kelas_kelasID = m.Kelas_kelasID
			datas.Role = m.Role
			datas.Status = sub.Status
			datas.TimeRemaining = &duration
			datas.Score = sub.Score
		}else if sub.Status == "submitted" {
			datas.ID = &sub.ID
			datas.Username = m.Username
			datas.User_userID = m.User_userID
			datas.Kelas_kelasID = m.Kelas_kelasID
			datas.Role = m.Role
			datas.Status = sub.Status
			datas.TimeRemaining = nil
			datas.Score = sub.Score
			}
		if sub.Status == entities.ExamStatus(flag){
		studentSubmission = append(studentSubmission, datas)
		}else if flag == "" {
			studentSubmission = append(studentSubmission, datas)
		}
	}
	return studentSubmission, nil
}

	// var res []dto.GetSubmissionStudentResponse
	// if flag == "" {
	// 	return result, nil
	// }
	
	// for _, m := range result {
	// 	if m.Status == entities.ExamStatus(flag){
	// 		res = append(res, m)
	// 	}
	// }


