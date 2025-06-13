package routes

import (
	controller "assesment/controller"
	"assesment/middleware"
	"assesment/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Answer(route *gin.Engine, injector *do.Injector) {
	answerController := do.MustInvoke[controller.AnswerController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, "jwtService")
	routes := route.Group("/answer")
	{
		// routes.GET("", answerController.GetAllAnswers)
		// routes.GET("/:id", answerController.GetAnswerByID)
		routes.POST("", middleware.Authenticate(jwtService), answerController.CreateAnswer)
		routes.PUT("/",middleware.Authenticate(jwtService), answerController.UpdateAnswer)
		routes.GET("/question/",middleware.Authenticate(jwtService), answerController.GetAnswerByQuestionID)
		routes.GET("/submission/",middleware.Authenticate(jwtService), answerController.ContinueSubmission)	
		// routes.GET("/student/:student_id", answerController.GetAnswerByStudentID)
	}
}


