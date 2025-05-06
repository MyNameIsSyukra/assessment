package routes

import (
	controller "assesment/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Question(route *gin.Engine, injector *do.Injector) {
	questionController := do.MustInvoke[controller.QuestionController](injector)

	routes := route.Group("assessment/question")
	{
		routes.POST("", questionController.CreateAllQuestion)
		routes.GET("", questionController.GetAllQuestions)
		routes.GET("/:id", questionController.GetQuestionByID)
		routes.PUT("/update", questionController.UpdateQuestion)
		routes.DELETE("/:id", questionController.DeleteQuestion)
	}
	routes = route.Group("assessment")
	{
		routes.GET("/detail/questions/:id", questionController.GetQuestionsByAssessmentID)
	}
}