package service

import (
	dto "assesment/dto"
	entities "assesment/entities"
	repository "assesment/repository"
	"assesment/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type (
	AssessmentService interface {
		// Teacher
		CreateAssessment(ctx context.Context,assesment *dto.AssessmentCreateRequest) (dto.AssessmentCreateResponse, error)
		TeacherGetAssessmentByID(ctx context.Context, id uuid.UUID) (dto.GetAssesmentByIDResponseTeacher, error)
		UpdateAssessment(ctx context.Context, assesment *dto.AssessmentUpdateRequest) (*entities.Assessment, error)
		DeleteAssessment(ctx context.Context, id uuid.UUID) error
		
		GetAllAssesmentByClassID(ctx context.Context, classID uuid.UUID) ([]entities.Assessment, error)
		// Student
		StudentGetAllAssesmentByClassIDAndUserID(ctx context.Context, classID uuid.UUID,userID uuid.UUID)([]dto.StudentGetAllAssesmentByClassIDResponse, error)
		GetAssessmentByIDAndUserID(ctx context.Context, classID uuid.UUID,userID uuid.UUID)(dto.GetAssessmentByIDAndByUserIDResponse, error)

		// Unused
		// GetAllAssessments(ctx context.Context)(dto.GetAllAssessmentsResponse, error)
	}
	assesmentService struct {
		assesmentRepo repository.AssessmentRepository
		submissionRepo repository.SubmissionRepository
	}
)

func NewAssessmentService(assesmentRepo repository.AssessmentRepository,submissionRepo repository.SubmissionRepository) AssessmentService {
	return &assesmentService{
		assesmentRepo: assesmentRepo,
		submissionRepo: submissionRepo,
	}
}

func (assesmentService *assesmentService) GetAllAssesmentByClassID(ctx context.Context, classID uuid.UUID) ([]entities.Assessment, error) {
	assessments, err := assesmentService.assesmentRepo.GetAllAssesmentByClassID(ctx, nil, classID)
	if err != nil {
		return []entities.Assessment{}, err
	}
	return assessments, nil
}

// teacherteacherteacherteacherteacherteacherteacherteacherteacherteacherteacherteacherteacherteacherteacherteacherteacherteacherteacherteacherteacherteacherteacher
func (assesmentService *assesmentService) CreateAssessment(ctx context.Context, assesment *dto.AssessmentCreateRequest) (dto.AssessmentCreateResponse, error) {
	// // checl if class is exist
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err)
	// }
	// params := url.Values{}
	// params.Add("id", assesment.ClassId.String())
	// urlClassSerivice := os.Getenv("CLASS_SERVICE_URL")
	// url := fmt.Sprintf("%s/kelas/?%s",urlClassSerivice,params.Encode())
	// resp, err := http.Get(url)
	// if err != nil {
	// 	return dto.AssessmentCreateResponse{}, fmt.Errorf("error checking class existence: %v", err)
	// }
	// if resp.StatusCode != 200 {
	// 	return dto.AssessmentCreateResponse{}, fmt.Errorf("no class found with id %s", assesment.ClassId.String())
	// }
	// defer resp.Body.Close()
	// // Baca body response
	
	assesmentEntity := entities.Assessment{
		ClassID: assesment.ClassId,
		Name: assesment.Name,
		Description: assesment.Description,
		Duration: assesment.Duration,
		CreatedAt: assesment.Date_created,
		StartTime: assesment.Start_time,
		EndTime: assesment.End_time,
	}
	
	createdAssessment, err := assesmentService.assesmentRepo.CreateAssessment(ctx, nil, &assesmentEntity)
	if err != nil {
		return dto.AssessmentCreateResponse{}, utils.ErrCreateAssesment
	}
	res := dto.AssessmentCreateResponse{
		ID: createdAssessment.ID,
		Name: createdAssessment.Name,
		Description: createdAssessment.Description,
		ClassId: createdAssessment.ClassID,
		Start_time: createdAssessment.StartTime,
		Duration: createdAssessment.Duration,
		End_time: createdAssessment.EndTime,
		Date_created: createdAssessment.CreatedAt,
		Updated_At: createdAssessment.UpdatedAt,
	}
	return res, nil	
}


func (assesmentService *assesmentService) TeacherGetAssessmentByID(ctx context.Context, id uuid.UUID) (dto.GetAssesmentByIDResponseTeacher,error){
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	urlClassSerivice := os.Getenv("CLASS_SERVICE_URL")
	assesment, err := assesmentService.assesmentRepo.GetAssessmentByID(ctx,nil,id)
	if err != nil {
		return dto.GetAssesmentByIDResponseTeacher{}, err
	}
	if assesment == nil {
		return dto.GetAssesmentByIDResponseTeacher{}, err
	}
	submission,err := assesmentService.submissionRepo.GetSubmissionsByAssessmentID(ctx, nil,id)
	if err != nil {
		return dto.GetAssesmentByIDResponseTeacher{}, err
	}
	if submission == nil {
		return dto.GetAssesmentByIDResponseTeacher{}, err
	}
	url := fmt.Sprintf("%s/service/class/%s",urlClassSerivice,assesment.ClassID)
	resp, err := http.Get(url)
	if err != nil {
		return dto.GetAssesmentByIDResponseTeacher{}, err
	}
	defer resp.Body.Close()
	// Baca body response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.GetAssesmentByIDResponseTeacher{}, err
	}
	// fmt.Print(body)
	// Unmarshal JSON ke struct
	var member []dto.GetMemberResponse
	err = json.Unmarshal(body, &member)
	if err != nil {
		return dto.GetAssesmentByIDResponseTeacher{}, err
	}
	// member := 
	totalSubmission := len(submission)
	totalStudent := len(member)
	return dto.GetAssesmentByIDResponseTeacher{
		ID: assesment.ID,
		Name: assesment.Name,
		Duration: assesment.Duration,
		StartTime: assesment.StartTime,
		EndTime: assesment.EndTime,
		TotalSubmission: totalSubmission,
		TotalStudent: totalStudent,
		// Questions: assesment.Questions,
	}, nil
}

func (assesmentService *assesmentService) UpdateAssessment(ctx context.Context, assesment *dto.AssessmentUpdateRequest) (*entities.Assessment, error) {
	ass,err := assesmentService.assesmentRepo.GetAssessmentByID(ctx, nil, assesment.Assessment_id)
	if ass == nil {
		return &entities.Assessment{}, errors.New("no assessment found")
	}
	if err != nil {
		return &entities.Assessment{}, err
	}
	if ass.StartTime.Before(time.Now()){
		return nil,errors.New("the assessment already started")
	}
	assesmentEntity := entities.Assessment{
		ID: ass.ID,
		ClassID: ass.ClassID,
		Description: assesment.Description,
		Name: assesment.Name,
		Duration: assesment.Duration,
		CreatedAt: assesment.Date_created,
		StartTime: assesment.Start_time,
		EndTime: assesment.End_time,
	}
	
	updatedAssessment, err := assesmentService.assesmentRepo.UpdateAssessment(ctx, nil, &assesmentEntity)
	if err != nil {
		return &entities.Assessment{}, err
	}
	return updatedAssessment, nil
}

func (assesmentService *assesmentService) DeleteAssessment(ctx context.Context, id uuid.UUID) error {
	asses,err := assesmentService.assesmentRepo.GetAssessmentByID(ctx, nil, id)
	if err != nil {
		return err
	}
	if asses == nil {
		return errors.New("no assessment found")
	}

	err = assesmentService.assesmentRepo.DeleteAssessment(ctx, nil,asses.ID.String())
	if err != nil {
		return err
	}
	return nil
}


// StudentStudentStudentStudentStudentStudentStudentStudentStudentStudentStudentStudentStudentStudentStudentStudentStudentStudentStudentStudentStudentStudent
func (assesmentService *assesmentService) StudentGetAllAssesmentByClassIDAndUserID(ctx context.Context, classID uuid.UUID,userID uuid.UUID)([]dto.StudentGetAllAssesmentByClassIDResponse, error){
	assessments, err := assesmentService.assesmentRepo.StudentGetAllAssesmentByClassIDAndUserID(ctx, nil, classID,userID)
	if err != nil {
		return []dto.StudentGetAllAssesmentByClassIDResponse{}, utils.ErrGetAllAssesmentByClassID
	}
	return assessments, nil
}

func (assesmentService *assesmentService) GetAssessmentByIDAndUserID(ctx context.Context, classID uuid.UUID,userID uuid.UUID)(dto.GetAssessmentByIDAndByUserIDResponse, error){
	assessments, err := assesmentService.assesmentRepo.GetAssessmentByIDAndByUserID(ctx, nil, classID,userID)
	if assessments == nil {
		return dto.GetAssessmentByIDAndByUserIDResponse{}, utils.ErrGetAssesmentByID
	}
	if err != nil {
		return dto.GetAssessmentByIDAndByUserIDResponse{}, utils.ErrGetAssesmentByID
	}

	return dto.GetAssessmentByIDAndByUserIDResponse{
		Assessment: assessments.Assessment,
		SubmittedAnswer: assessments.SubmittedAnswer,
		Question: assessments.Question,
		TimeSpent: assessments.TimeSpent,
		TimeRemaining: assessments.TimeRemaining,
		MaxScore: assessments.MaxScore,
		Score: assessments.Score,
		SubmissionStatus: assessments.SubmissionStatus,
		SubmissionID: assessments.SubmissionID,
	}, nil 
}

// Antar Service






















// func (assesmentService *assesmentService) GetAllAssessments (ctx context.Context) (dto.GetAllAssessmentsResponse, error) {
// 	assessments, err := assesmentService.assesmentRepo.GetAllAssessments()
// 	if len(assessments) == 0 {
// 		return dto.GetAllAssessmentsResponse{}, utils.ErrGetAllAssesments
// 	}
// 	if err != nil {
// 		return dto.GetAllAssessmentsResponse{}, err
// 	}
// 	return dto.GetAllAssessmentsResponse{
// 		Assessments: assessments,
// 	}, nil
// }