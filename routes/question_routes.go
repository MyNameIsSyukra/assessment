package routes

import (
	controller "assesment/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Question(route *gin.Engine, injector *do.Injector) {
	questionController := do.MustInvoke[controller.QuestionController](injector)

	routes := route.Group("/api/v1/question")
	{
		routes.POST("", questionController.CreateQuestion)
		routes.GET("", questionController.GetAllQuestions)
		routes.GET("/:id", questionController.GetQuestionByID)
		routes.PUT("/:id", questionController.UpdateQuestion)
		routes.DELETE("/:id", questionController.DeleteQuestion)
		routes.GET("/assessment/:assessment_id", questionController.GetQuestionsByAssessmentID)
	}
}