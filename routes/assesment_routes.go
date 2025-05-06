package routes

import (
	controller "assesment/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Assessment(route *gin.Engine, injector *do.Injector) {
	asssessmentController := do.MustInvoke[controller.AssessmentController](injector)

	routes := route.Group("teacher/assessment")
	{
		routes.POST("", asssessmentController.CreateAssessment)
		routes.GET("/:id", asssessmentController.TeacherGetAssessmentByID)
		routes.PUT("/:id", asssessmentController.UpdateAssessment)
		routes.DELETE("/:id", asssessmentController.DeleteAssessment)
		routes.GET("/class/:classID", asssessmentController.GetAllAssesmentByClassID)
		// routes.GET("", asssessmentController.GetAllAssessments)
	}

	// Student routes
	routes = route.Group("/student/assessment")
	{
		routes.GET("/:id/:userID", asssessmentController.GetAssessmentByIDAndUserID)
		routes.GET("/class/:classID/:userID", asssessmentController.StudentGetAllAssesmentByClassIDAssesmentFlag)
	}
}