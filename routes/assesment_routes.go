package routes

import (
	controller "assesment/controller"
	"assesment/middleware"
	"assesment/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Assessment(route *gin.Engine, injector *do.Injector) {
	asssessmentController := do.MustInvoke[controller.AssessmentController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, "jwtService")
	routes := route.Group("teacher/assessment")
	{
		routes.POST("",middleware.Authenticate(jwtService),middleware.RequireTeacherRole(jwtService), asssessmentController.CreateAssessment)
		routes.GET("",middleware.Authenticate(jwtService),middleware.RequireTeacherRole(jwtService), asssessmentController.TeacherGetAssessmentByID)
		routes.PUT("/update",middleware.Authenticate(jwtService),middleware.RequireTeacherRole(jwtService), asssessmentController.UpdateAssessment)
		routes.DELETE("/delete",middleware.Authenticate(jwtService),middleware.RequireTeacherRole(jwtService), asssessmentController.DeleteAssessment)
		routes.GET("/class/", middleware.Authenticate(jwtService),middleware.RequireTeacherRole(jwtService),asssessmentController.GetAllAssesmentByClassID)
	}

	// Student routes
	routes = route.Group("/student/assessment")
	{
		routes.GET("/",middleware.Authenticate(jwtService), asssessmentController.GetAssessmentByIDAndUserID)
		routes.GET("/class/", middleware.Authenticate(jwtService),asssessmentController.StudentGetAllAssesmentByClassIDAssesmentFlag)
	}
}