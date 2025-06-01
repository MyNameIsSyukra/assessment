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
		routes.GET("", asssessmentController.TeacherGetAssessmentByID)
		routes.PUT("/update", asssessmentController.UpdateAssessment)
		routes.DELETE("/delete", asssessmentController.DeleteAssessment)
		routes.GET("/class/", asssessmentController.GetAllAssesmentByClassID)
	}

	// Student routes
	routes = route.Group("/student/assessment")
	{
		routes.GET("/", asssessmentController.GetAssessmentByIDAndUserID)
		routes.GET("/class/", asssessmentController.StudentGetAllAssesmentByClassIDAssesmentFlag)
	}
}