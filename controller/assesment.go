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
	AssessmentController interface {
		// Teacher
		CreateAssessment(ctx *gin.Context)
		TeacherGetAssessmentByID(ctx *gin.Context) 	
		UpdateAssessment(ctx *gin.Context) 
		DeleteAssessment(ctx *gin.Context) 
		GetAllAssesmentByClassID(ctx *gin.Context)

		// Student
		StudentGetAllAssesmentByClassIDAssesmentFlag(ctx *gin.Context)
		GetAssessmentByIDAndUserID(ctx *gin.Context)

		// Lintas Service
		ServiceGetAllAssesmentByClassIDAssesmentFlag(ctx *gin.Context)
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

func (assesmentController *assesmentController) GetAllAssesmentByClassID(ctx *gin.Context) {
	classID,err := uuid.Parse(ctx.Query("classID"))
	if err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	assessments, err := assesmentController.assesmentService.GetAllAssesmentByClassID(ctx.Request.Context(), classID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(assessments)
	ctx.JSON(http.StatusOK, res)
}

// ===========================================Teacher==================================================
func (assesmentController *assesmentController) TeacherGetAssessmentByID(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Query("id"))
	if err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	assesment, err := assesmentController.assesmentService.TeacherGetAssessmentByID(ctx.Request.Context(), id)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(assesment)
	ctx.JSON(http.StatusOK, res)
}

func (assesmentController *assesmentController) CreateAssessment(ctx *gin.Context) {
	var request dto.AssessmentCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	assesment, err := assesmentController.assesmentService.CreateAssessment(ctx.Request.Context(), &request)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(assesment)
	ctx.JSON(http.StatusCreated, res)
}

func (assesmentController *assesmentController) UpdateAssessment(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Query("id"))
	if err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var request dto.AssessmentUpdateRequest
	request.IdEvaluation = id
	if err := ctx.ShouldBindJSON(&request); err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	assesment, err := assesmentController.assesmentService.UpdateAssessment(ctx.Request.Context(), &request)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(assesment)
	ctx.JSON(http.StatusOK, res)
}

func (assesmentController *assesmentController) DeleteAssessment(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Query("id"))
	if err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = assesmentController.assesmentService.DeleteAssessment(ctx.Request.Context(), id)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse("Deleted successfully")
	ctx.JSON(http.StatusOK, res)
}


// ==========================================Student==================================================
func (assesmentController *assesmentController) GetAssessmentByIDAndUserID(ctx *gin.Context) {
	id,err := uuid.Parse(ctx.Query("id"))
	if err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	userID,err := uuid.Parse(ctx.Query("userID"))
	if err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	assesment, err := assesmentController.assesmentService.GetAssessmentByIDAndUserID(ctx.Request.Context(), id,userID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(assesment)
	ctx.JSON(http.StatusOK, res)
}

func (assesmentController *assesmentController) StudentGetAllAssesmentByClassIDAssesmentFlag(ctx *gin.Context) {
	classID,err := uuid.Parse(ctx.Query("classID"))
	if err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	userID,err := uuid.Parse(ctx.Query("userID"))
	if err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	assessments, err := assesmentController.assesmentService.StudentGetAllAssesmentByClassIDAndUserID(ctx.Request.Context(), classID,userID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.SuccessResponse(assessments)
	ctx.JSON(http.StatusOK, res)
}


//=================================== Lintas Service======================================================
func (assesmentController *assesmentController) ServiceGetAllAssesmentByClassIDAssesmentFlag(ctx *gin.Context) {
	classID,err := uuid.Parse(ctx.Param("classID"))
	if err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	userID,err := uuid.Parse(ctx.Param("userID"))
	if err != nil {
		res := utils.FailedResponse(utils.FailedGetDataFromBody)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	assessments, err := assesmentController.assesmentService.StudentGetAllAssesmentByClassIDAndUserID(ctx.Request.Context(), classID,userID)
	if err != nil {
		res := utils.FailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	// res := utils.SuccessResponse(assessments)
	ctx.JSON(http.StatusOK, assessments)
}









// Unused code
// GetAllAssessments(ctx *gin.Context) 
// func (assesmentController *assesmentController) GetAllAssessments(ctx *gin.Context) {
// 	assessments, err := assesmentController.assesmentService.GetAllAssessments(ctx.Request.Context())
// 	if err != nil {
// 		res := utils.FailedResponse(err.Error())
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}
// 	res := utils.SuccessResponse(assessments)
// 	ctx.JSON(http.StatusOK, res)
// }