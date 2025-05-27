package routes

import (
	controller "assesment/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Answer(route *gin.Engine, injector *do.Injector) {
	answerController := do.MustInvoke[controller.AnswerController](injector)

	routes := route.Group("/answer")
	{
		// routes.GET("", answerController.GetAllAnswers)
		// routes.GET("/:id", answerController.GetAnswerByID)
		routes.POST("", answerController.CreateAnswer)
		routes.PUT("/", answerController.UpdateAnswer)
		routes.GET("/question/", answerController.GetAnswerByQuestionID)
		routes.GET("/submission/", answerController.GetAnswerBySubmissionID)	
		// routes.GET("/student/:student_id", answerController.GetAnswerByStudentID)
	}
}


