package routes

import (
	controller "assesment/controller"
	"assesment/middleware"
	"assesment/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Submission(route *gin.Engine, injector *do.Injector) {
	submissionController := do.MustInvoke[controller.SubmissionController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, "jwtService")
	routes := route.Group("/submission")
	{
		routes.POST("",middleware.Authenticate(jwtService), submissionController.CreateSubmission)
		// routes.GET("", submissionController.GetAllSubmissions)
		routes.GET("/",middleware.Authenticate(jwtService), submissionController.GetSubmissionByID)
		// routes.GET("/user/", submissionController.GetSubmissionsByUserID)
		routes.POST("/submit/",middleware.Authenticate(jwtService), submissionController.Submitted)
	}
	
	// teacher
	routes = route.Group("assement/submission")
	{
		routes.GET("/",middleware.Authenticate(jwtService),middleware.RequireTeacherRole(jwtService),submissionController.GetStudentSubmissionsByAssessmentID)
		routes.DELETE("/",middleware.Authenticate(jwtService),middleware.RequireTeacherRole(jwtService), submissionController.DeleteSubmission)
	}
}

// routes.GET("/assessment/:assessment_id/:assessment_id/:user_id", submissionController.GetSubmissionsByAssessmentIDAndUserID)
// routes.GET("/assessment/:assessment_id/class/:class_id/:assessment_id",submissionController.GetSubmissionsByAssessmentIDAndClassID)