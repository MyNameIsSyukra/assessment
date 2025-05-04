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
	QuestionController interface {
		CreateQuestion(ctx *gin.Context)
		GetAllQuestions(ctx *gin.Context)
		GetQuestionByID(ctx *gin.Context)
		UpdateQuestion(ctx *gin.Context)
		DeleteQuestion(ctx *gin.Context)
		GetQuestionsByAssessmentID(ctx *gin.Context)
		CreateAllQuestion(ctx *gin.Context)
	}
	questionController struct {
		questionService service.QuestionService
	}
)

func NewQuestionController(questionService service.QuestionService) QuestionController {
	return &questionController{
		questionService: questionService,
	}
}


func (questionController *questionController) CreateAllQuestion(ctx *gin.Context) {
	var request dto.CreateAllQuestionRequest
	var response []dto.QuestionResponse
	if err := ctx.ShouldBindJSON(&request); err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	for _, choice := range request.Questions {
		data, err := questionController.questionService.CreatePertanyaan(ctx.Request.Context(), request.EvaluationID, choice)
		if err != nil {
			res := utils.FailedResponse(err.Error())
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
		response = append(response, data)	
	}
	res := utils.SuccessResponse(response)
	ctx.JSON(http.StatusCreated, res)
}


func (questionController *questionController) CreateQuestion(ctx *gin.Context) {
	var request dto.QuestionCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	question, err := questionController.questionService.CreateQuestion(ctx.Request.Context(), &request)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(question)
	ctx.JSON(http.StatusCreated, res)
}	

func (questionController *questionController) GetAllQuestions(ctx *gin.Context) {
	questions, err := questionController.questionService.GetAllQuestions(ctx.Request.Context())
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(questions)
	ctx.JSON(http.StatusOK, res)
}

func (questionController *questionController) GetQuestionByID(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.FailedResponse("invalid id format")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	question, err := questionController.questionService.GetQuestionByID(ctx.Request.Context(), id)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(question)
	ctx.JSON(http.StatusOK, res)
}

func (questionController *questionController) UpdateQuestion(ctx *gin.Context) {
	var request dto.QuestionUpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	// fmt.Print(request)
	question, err := questionController.questionService.UpdateQuestion(ctx.Request.Context(), &request)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(question)
	ctx.JSON(http.StatusOK, res)
}

func (questionController *questionController) DeleteQuestion(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.FailedResponse("invalid id format")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	err = questionController.questionService.DeleteQuestion(ctx.Request.Context(), id)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse("Deleted successfully")
	ctx.JSON(http.StatusOK, res)
}

func (questionController *questionController) GetQuestionsByAssessmentID(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.FailedResponse("invalid id format")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	questions, err := questionController.questionService.GetQuestionsByAssessmentID(ctx.Request.Context(), id)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(questions)
	ctx.JSON(http.StatusOK, res)
}
