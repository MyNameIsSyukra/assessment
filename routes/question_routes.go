package routes

import (
	controller "assesment/controller"
	"assesment/middleware"
	"assesment/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Question(route *gin.Engine, injector *do.Injector) {
	questionController := do.MustInvoke[controller.QuestionController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, "jwtService")
	// Teacher
	routes := route.Group("assessment/question")
	{
		routes.POST("",middleware.Authenticate(jwtService), questionController.CreateAllQuestion)
		// routes.GET("", questionController.GetAllQuestions)
		routes.GET("/",middleware.Authenticate(jwtService), questionController.GetQuestionByID)
		routes.PUT("/update",middleware.Authenticate(jwtService), questionController.UpdateQuestion)
		routes.DELETE("/",middleware.Authenticate(jwtService), questionController.DeleteQuestion)
	}

	// Student
	routes = route.Group("assessment")
	{
		routes.GET("/detail/questions/",middleware.Authenticate(jwtService), questionController.GetQuestionsByAssessmentID)
	}
}