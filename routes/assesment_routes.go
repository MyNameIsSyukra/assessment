package routes

import (
	controller "assesment/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Assessment(route *gin.Engine, injector *do.Injector) {
	asssessmentController := do.MustInvoke[controller.AssessmentController](injector)

	routes := route.Group("/api/v1/assessment")
	{
		routes.POST("", asssessmentController.CreateAssessment)
		routes.GET("", asssessmentController.GetAllAssessments)
		routes.GET("/:id", asssessmentController.GetAssessmentByID)
		routes.PUT("/:id", asssessmentController.UpdateAssessment)
		routes.DELETE("/:id", asssessmentController.DeleteAssessment)
		routes.GET("/class/:classID", asssessmentController.GetAllAssesmentByClassID)
	}

	// Student routes
	routes = route.Group("/api/v1/student/assessment")
	{
		routes.GET("/class/:classID/:userID", asssessmentController.GetAllAssesmentByClassIDAssesmentFlag)
		routes.GET("/:id/:userID", asssessmentController.GetAssessmentByIDAndUserID)
	}

}