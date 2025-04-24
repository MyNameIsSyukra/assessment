package controller

import (
	dto "assesment/dto"
	service "assesment/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	QuestionController interface {
		CreateQuestion(ctx *gin.Context)
		GetAllQuestions(ctx *gin.Context)
		GetQuestionByID(ctx *gin.Context)
		UpdateQuestion(ctx *gin.Context)
		DeleteQuestion(ctx *gin.Context)
		GetQuestionsByAssessmentID(ctx *gin.Context)
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

func (questionController *questionController) CreateQuestion(ctx *gin.Context) {
	var request dto.QuestionCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question, err := questionController.questionService.CreateQuestion(ctx.Request.Context(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, question)
}	

func (questionController *questionController) GetAllQuestions(ctx *gin.Context) {
	questions, err := questionController.questionService.GetAllQuestions(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, questions)
}

func (questionController *questionController) GetQuestionByID(ctx *gin.Context) {
	id := ctx.Param("id")
	question, err := questionController.questionService.GetQuestionByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, question)
}

func (questionController *questionController) UpdateQuestion(ctx *gin.Context) {
	id := ctx.Param("id")
	var request dto.QuestionUpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request.Id = id

	question, err := questionController.questionService.UpdateQuestion(ctx.Request.Context(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, question)
}

func (questionController *questionController) DeleteQuestion(ctx *gin.Context) {
	id := ctx.Param("id")
	err := questionController.questionService.DeleteQuestion(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (questionController *questionController) GetQuestionsByAssessmentID(ctx *gin.Context) {
	id := ctx.Param("id")
	questions, err := questionController.questionService.GetQuestionsByAssessmentID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, questions)
}
