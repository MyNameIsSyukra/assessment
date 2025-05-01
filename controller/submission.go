package controller

import (
	dto "assesment/dto"
	service "assesment/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (	
	SubmissionController interface {
		CreateSubmission(ctx *gin.Context)
		GetAllSubmissions(ctx *gin.Context)
		GetSubmissionByID(ctx *gin.Context)
		// UpdateSubmission(ctx *gin.Context)
		DeleteSubmission(ctx *gin.Context)
		GetSubmissionsByAssessmentID(ctx *gin.Context)
		GetSubmissionsByUserID(ctx *gin.Context)
		GetSubmissionsByAssessmentIDAndUserID(ctx *gin.Context)
		GetSubmissionsByAssessmentIDAndClassID(ctx *gin.Context)
	}

	submissionController struct {
		submissionService service.SubmissionService
	}
)

func NewSubmissionController(submissionService service.SubmissionService) SubmissionController {
	return &submissionController{
		submissionService: submissionService,
	}
}

func (submissionController *submissionController) CreateSubmission(ctx *gin.Context) {
	var request dto.SubmissionCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	submission, err := submissionController.submissionService.CreateSubmission(ctx.Request.Context(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, submission)
}

func (submissionController *submissionController) GetAllSubmissions(ctx *gin.Context) {
	submissions, err := submissionController.submissionService.GetAllSubmissions(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, submissions)
}

func (submissionController *submissionController) GetSubmissionByID(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	submission, err := submissionController.submissionService.GetSubmissionByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, submission)
}
// func (submissionController *submissionController) UpdateSubmission(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	var request dto.SubmissionUpdateRequest
// 	if err := ctx.ShouldBindJSON(&request); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	request.Id = id
// 	submission, err := submissionController.submissionService.UpdateSubmission(ctx.Request.Context(), &request)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, submission)
// }

func (submissionController *submissionController) DeleteSubmission(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	err = submissionController.submissionService.DeleteSubmission(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (submissionController *submissionController) GetSubmissionsByAssessmentID(ctx *gin.Context) {
	assessmentID,err := uuid.Parse(ctx.Param("assessment_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	submissions, err := submissionController.submissionService.GetSubmissionsByAssessmentID(ctx.Request.Context(), assessmentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, submissions)
}

func (submissionController *submissionController) GetSubmissionsByUserID(ctx *gin.Context) {
	userID,err := uuid.Parse(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	submissions, err := submissionController.submissionService.GetSubmissionsByUserID(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, submissions)
}

func (submissionController *submissionController) GetSubmissionsByAssessmentIDAndUserID(ctx *gin.Context) {
	assessmentID,err := uuid.Parse(ctx.Param("assessment_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	userID,err := uuid.Parse(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	submission, err := submissionController.submissionService.GetSubmissionsByAssessmentIDAndUserID(ctx.Request.Context(), assessmentID, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, submission)
}

func (submissionController *submissionController) GetSubmissionsByAssessmentIDAndClassID(ctx *gin.Context) {
	assessmentID,err := uuid.Parse(ctx.Param("assessment_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	classID,err := uuid.Parse(ctx.Param("class_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	submissions, err := submissionController.submissionService.GetSubmissionsByAssessmentIDAndClassID(ctx.Request.Context(), assessmentID, classID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, submissions)
}

