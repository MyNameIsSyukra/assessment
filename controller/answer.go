package controller

import (
	dto "assesment/dto"
	service "assesment/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	AnswerController interface {
		CreateAnswer(ctx *gin.Context)
		GetAllAnswers(ctx *gin.Context)
		GetAnswerByID(ctx *gin.Context)
		UpdateAnswer(ctx *gin.Context)
		GetAnswerByQuestionID(ctx *gin.Context)
		GetAnswerByStudentID(ctx *gin.Context)
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	answer, err := answerController.answerService.CreateAnswer(ctx.Request.Context(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, answer)
}

func (answerController *answerController) GetAllAnswers(ctx *gin.Context) {
	answers, err := answerController.answerService.GetAllAnswers(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, answers)
}

func (answerController *answerController) GetAnswerByID(ctx *gin.Context) {
	id := ctx.Param("id")
	answer, err := answerController.answerService.GetAnswerByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, answer)
}

func (answerController *answerController) UpdateAnswer(ctx *gin.Context) {
	id := ctx.Param("id")
	var request dto.AnswerUpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request.ID = id
	answer, err := answerController.answerService.UpdateAnswer(ctx.Request.Context(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, answer)
}

func (answerController *answerController) GetAnswerByQuestionID(ctx *gin.Context) {
	questionID := ctx.Param("question_id")
	answers, err := answerController.answerService.GetAnswerByQuestionID(ctx.Request.Context(), questionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, answers)
}

func (answerController *answerController) GetAnswerByStudentID(ctx *gin.Context) {
	studentID := ctx.Param("student_id")
	answers, err := answerController.answerService.GetAnswerByStudentID(ctx.Request.Context(), studentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, answers)
}

