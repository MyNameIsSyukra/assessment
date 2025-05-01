package controller

import (
	dto "assesment/dto"
	service "assesment/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	AssessmentController interface {
		CreateAssessment(ctx *gin.Context)
		GetAllAssessments(ctx *gin.Context) 
		GetAssessmentByID(ctx *gin.Context) 	
		UpdateAssessment(ctx *gin.Context) 
		DeleteAssessment(ctx *gin.Context) 
		GetAllAssesmentByClassID(ctx *gin.Context)
	}
	assesmentController struct {
	assesmentService service.AssessmentService
	}
	)

func NewAssessmentController(assesmentService service.AssessmentService) AssessmentController {
	return &assesmentController{
		assesmentService: assesmentService,
	}
}

func (assesmentController *assesmentController) CreateAssessment(ctx *gin.Context) {
	var request dto.AssessmentCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.FailedGetDataFromBody})
		return
	}

	assesment, err := assesmentController.assesmentService.CreateAssessment(ctx.Request.Context(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, assesment)
}

func (assesmentController *assesmentController) GetAllAssessments(ctx *gin.Context) {
	assessments, err := assesmentController.assesmentService.GetAllAssessments(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, assessments)
}

func (assesmentController *assesmentController) GetAllAssesmentByClassID(ctx *gin.Context) {
	classID,err := uuid.Parse(ctx.Param("classID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}
	assessments, err := assesmentController.assesmentService.GetAllAssesmentByClassID(ctx.Request.Context(), classID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, assessments)
}

func (assesmentController *assesmentController) GetAssessmentByID(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	assesment, err := assesmentController.assesmentService.GetAssessmentByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, assesment)
}

func (assesmentController *assesmentController) UpdateAssessment(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var request dto.AssessmentUpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request.IdEvaluation = id

	assesment, err := assesmentController.assesmentService.UpdateAssessment(ctx.Request.Context(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, assesment)
}

func (assesmentController *assesmentController) DeleteAssessment(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = assesmentController.assesmentService.DeleteAssessment(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Deleted successfully")
}

