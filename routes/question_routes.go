package routes

import (
	controller "assesment/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Question(route *gin.Engine, injector *do.Injector) {
	questionController := do.MustInvoke[controller.QuestionController](injector)

	// Teacher
	routes := route.Group("assessment/question")
	{
		routes.POST("", questionController.CreateAllQuestion)
		routes.GET("", questionController.GetAllQuestions)
		routes.GET("/", questionController.GetQuestionByID)
		routes.PUT("/update", questionController.UpdateQuestion)
		routes.DELETE("/", questionController.DeleteQuestion)
	}

	// Student
	routes = route.Group("assessment")
	{
		routes.GET("/detail/questions/", questionController.GetQuestionsByAssessmentID)
	}
}