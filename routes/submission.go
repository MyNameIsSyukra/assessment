package routes

import (
	controller "assesment/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Submission(route *gin.Engine, injector *do.Injector) {
	submissionController := do.MustInvoke[controller.SubmissionController](injector)

	routes := route.Group("/submission")
	{
		routes.POST("", submissionController.CreateSubmission)
		// routes.GET("", submissionController.GetAllSubmissions)
		routes.GET("/", submissionController.GetSubmissionByID)
		routes.DELETE("/", submissionController.DeleteSubmission)
		routes.GET("/user/", submissionController.GetSubmissionsByUserID)
		routes.POST("/submit/", submissionController.Submitted)
	}
	
	// teacher
	routes = route.Group("assement/submission")
	{
		routes.GET("/",submissionController.GetStudentSubmissionsByAssessmentID)
	}
}

// routes.GET("/assessment/:assessment_id/:assessment_id/:user_id", submissionController.GetSubmissionsByAssessmentIDAndUserID)
// routes.GET("/assessment/:assessment_id/class/:class_id/:assessment_id",submissionController.GetSubmissionsByAssessmentIDAndClassID)