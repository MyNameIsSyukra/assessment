package controller

import (
	dto "assesment/dto"
	service "assesment/service"
	"assesment/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	AnswerController interface {
		CreateAnswer(ctx *gin.Context)
		GetAllAnswers(ctx *gin.Context)
		GetAnswerByID(ctx *gin.Context)
		UpdateAnswer(ctx *gin.Context)
		GetAnswerByQuestionID(ctx *gin.Context)
		GetAnswerBySubmissionID(ctx *gin.Context)
		// GetAnswerByStudentID(ctx *gin.Context)
	}
	answerController struct {
		answerService service.AnswerService
	}
)

func NewAnswerController(answerService service.AnswerService) AnswerController {
	return &answerController{
		answerService: answerService,
	}
}

func (answerController *answerController) CreateAnswer(ctx *gin.Context) {
	var request dto.AnswerCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	answer, err := answerController.answerService.CreateAnswer(ctx.Request.Context(), &request)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(answer)
	ctx.JSON(http.StatusCreated, res)
}

func (answerController *answerController) GetAllAnswers(ctx *gin.Context) {
	answers, err := answerController.answerService.GetAllAnswers(ctx.Request.Context())
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(answers)
	ctx.JSON(http.StatusOK, res)
}

func (answerController *answerController) GetAnswerByID(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	answer, err := answerController.answerService.GetAnswerByID(ctx.Request.Context(), id)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(answer)
	ctx.JSON(http.StatusOK, res)
}

func (answerController *answerController) UpdateAnswer(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	var request dto.AnswerUpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	request.ID = id
	answer, err := answerController.answerService.UpdateAnswer(ctx.Request.Context(), &request)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(answer)
	ctx.JSON(http.StatusOK, res)
}

func (answerController *answerController) GetAnswerByQuestionID(ctx *gin.Context) {
	questionID,err := uuid.Parse(ctx.Param("question_id"))
	if err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	
	answers, err := answerController.answerService.GetAnswerByQuestionID(ctx.Request.Context(), questionID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	ctx.JSON(http.StatusOK, answers)
}

func (answerController *answerController) GetAnswerBySubmissionID(ctx *gin.Context) {
	submissionID, err := uuid.Parse(ctx.Param("submission_id"))
	if err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	answers, err := answerController.answerService.GetAnswerBySubmissionID(ctx.Request.Context(), submissionID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	ctx.JSON(http.StatusOK, answers)
}



// func (answerController *answerController) GetAnswerByStudentID(ctx *gin.Context) {
// 	var req dto.GetAnswerByStudentIDRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.FailedGetDataFromBody})
// 		return
// 	}
// 	answers, err := answerController.answerService.GetAnswerByStudentID(ctx.Request.Context(), req)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, answers)
// }

