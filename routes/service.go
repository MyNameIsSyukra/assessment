package routes

import (
	controller "assesment/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Service(route *gin.Engine, injector *do.Injector) {
	// questionController := do.MustInvoke[controller.QuestionController](injector)
	assesmentController := do.MustInvoke[controller.AssessmentController](injector)
	// answerController := do.MustInvoke[controller.AnswerController](injector)
	// submissionController := do.MustInvoke[controller.SubmissionController](injector)

	routes := route.Group("service")
	{
		routes.GET("assessment/class/:classID/:userID", assesmentController.ServiceGetAllAssesmentByClassIDAssesmentFlag)
	}
}