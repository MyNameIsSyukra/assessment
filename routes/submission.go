package routes

import (
	controller "assesment/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Submission(route *gin.Engine, injector *do.Injector) {
	submissionController := do.MustInvoke[controller.SubmissionController](injector)

	routes := route.Group("/api/v1/submission")
	{
		routes.POST("", submissionController.CreateSubmission)
		routes.GET("", submissionController.GetAllSubmissions)
		routes.GET("/:id", submissionController.GetSubmissionByID)
		routes.DELETE("/:id", submissionController.DeleteSubmission)
		routes.GET("/assessment/:assessment_id", submissionController.GetSubmissionsByAssessmentID)
		routes.GET("/user/:user_id", submissionController.GetSubmissionsByUserID)
		routes.GET("/assessment/:assessment_id/:assessment_id/:user_id", submissionController.GetSubmissionsByAssessmentIDAndUserID)
		routes.GET("/assessment/:assessment_id/class/:class_id/:assessment_id",submissionController.GetSubmissionsByAssessmentIDAndClassID)
		routes.POST("/submit/:id", submissionController.Submitted)
	}
}
