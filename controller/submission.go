package controller

import (
	dto "assesment/dto"
	service "assesment/service"
	"assesment/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (	
	SubmissionController interface {
		CreateSubmission(ctx *gin.Context)
		GetAllSubmissions(ctx *gin.Context)
		GetSubmissionByID(ctx *gin.Context)
		DeleteSubmission(ctx *gin.Context)
		GetSubmissionsByUserID(ctx *gin.Context)
		// GetSubmissionsByAssessmentIDAndUserID(ctx *gin.Context)
		GetStudentSubmissionsByAssessmentID(ctx *gin.Context)
		Submitted(ctx *gin.Context)
		// UpdateSubmission(ctx *gin.Context)
		// GetSubmissionsByAssessmentID(ctx *gin.Context)
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
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	submission, err := submissionController.submissionService.CreateSubmission(ctx.Request.Context(), &request)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	ctx.JSON(http.StatusCreated, submission)
}

func (submissionController *submissionController) GetAllSubmissions(ctx *gin.Context) {
	submissions, err := submissionController.submissionService.GetAllSubmissions(ctx.Request.Context())
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(submissions)
	ctx.JSON(http.StatusOK, res)
}

func (submissionController *submissionController) GetSubmissionByID(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Query("id"))
	if err != nil {
		res := utils.FailedResponse("invalid id format")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	submission, err := submissionController.submissionService.GetSubmissionByID(ctx.Request.Context(), id)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(submission)
	ctx.JSON(http.StatusOK, res)
}


func (submissionController *submissionController) DeleteSubmission(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Query("id"))
	if err != nil {
		res := utils.FailedResponse("invalid id format")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	err = submissionController.submissionService.DeleteSubmission(ctx.Request.Context(), id)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse("submission deleted successfully")
	ctx.JSON(http.StatusNoContent, res)
}

func (submissionController *submissionController) GetSubmissionsByUserID(ctx *gin.Context) {
	userID,err := uuid.Parse(ctx.Query("user_id"))
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
	assessmentID,err := uuid.Parse(ctx.Query("assessment_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	userID,err := uuid.Parse(ctx.Query("user_id"))
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

func (submissionController *submissionController) Submitted(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	submission, err := submissionController.submissionService.Submitted(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, submission)
}

func (submissionController *submissionController)GetStudentSubmissionsByAssessmentID(ctx *gin.Context)(){
	id,err := uuid.Parse(ctx.Query("assessment_id"))
	flag := ctx.Query("status")
	fmt.Println(flag)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	submissions, err := submissionController.submissionService.GetStudentSubmissionsByAssessmentID(ctx.Request.Context(),id,flag)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest,res)
		return
	}
	res := utils.SuccessResponse(submissions)
	ctx.JSON(http.StatusOK,res)
}


// func (submissionController *submissionController) GetSubmissionsByAssessmentID(ctx *gin.Context) {
// 	assessmentID,err := uuid.Parse(ctx.Param("assessment_id"))
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}
// 	submissions, err := submissionController.submissionService.GetSubmissionsByAssessmentID(ctx.Request.Context(), assessmentID)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, submissions)
// }

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