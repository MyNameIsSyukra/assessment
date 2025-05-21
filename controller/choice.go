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
	ChoiceController interface {
		CreateChoice(ctx *gin.Context)
		GetChoiceByID(ctx *gin.Context)
		// UpdateChoice(ctx *gin.Context)
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
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	choice, err := choiceController.choiceService.CreateChoice(ctx.Request.Context(), &request)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(choice)
	ctx.JSON(http.StatusCreated, res)
}

func (choiceController *choiceController) GetChoiceByID(ctx *gin.Context) {
	id ,err:= uuid.Parse(ctx.Query("id"))
	if err != nil {
		res := utils.FailedResponse("invalid id format")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	choice, err := choiceController.choiceService.GetChoiceByID(ctx.Request.Context(), id)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(choice)
	ctx.JSON(http.StatusOK, res)
}

// func (choiceController *choiceController) UpdateChoice(ctx *gin.Context) {
// 	id,err:= uuid.Parse(ctx.Param("id"))
// 	if err != nil {
// 		res := utils.FailedResponse("invalid id format")
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}
// 	var request dto.ChoiceUpdateRequest
// 	request.ID = id
// 	if err := ctx.ShouldBindJSON(&request); err != nil {
// 		res := utils.FailedResponse(utils.FailedGetDataFromBody)
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}
// 	choice, err := choiceController.choiceService.UpdateChoice(ctx.Request.Context(), &request)
// 	if err != nil {
// 		res := utils.FailedResponse(err.Error())
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}
// 	res := utils.SuccessResponse(choice)
// 	ctx.JSON(http.StatusOK, res)
// }

func (choiceController *choiceController) GetChoicesByQuestionID(ctx *gin.Context) {
	questionID,err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.FailedResponse("invalid id format")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	choices, err := choiceController.choiceService.GetChoiceByQuestionID(ctx.Request.Context(), questionID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(choices)
	ctx.JSON(http.StatusOK, res)
}