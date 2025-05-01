package controller

import (
	dto "assesment/dto"
	service "assesment/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	ChoiceController interface {
		CreateChoice(ctx *gin.Context)
		GetChoiceByID(ctx *gin.Context)
		UpdateChoice(ctx *gin.Context)
		GetChoicesByQuestionID(ctx *gin.Context)
	}
	choiceController struct {
		choiceService service.ChoiceService
	}
)

func NewChoiceController(choiceService service.ChoiceService) ChoiceController {
	return &choiceController{
		choiceService: choiceService,
	}
}

func (choiceController *choiceController) CreateChoice(ctx *gin.Context) {
	var request dto.ChoiceCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.FailedGetDataFromBody})
		return
	}

	choice, err := choiceController.choiceService.CreateChoice(ctx.Request.Context(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, choice)
}

func (choiceController *choiceController) GetChoiceByID(ctx *gin.Context) {
	id ,err:= uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	choice, err := choiceController.choiceService.GetChoiceByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, choice)
}

func (choiceController *choiceController) UpdateChoice(ctx *gin.Context) {
	id,err:= uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	var request dto.ChoiceUpdateRequest
	request.ID = id
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	choice, err := choiceController.choiceService.UpdateChoice(ctx.Request.Context(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, choice)
}

func (choiceController *choiceController) GetChoicesByQuestionID(ctx *gin.Context) {
	questionID,err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	choices, err := choiceController.choiceService.GetChoiceByQuestionID(ctx.Request.Context(), questionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, choices)
}