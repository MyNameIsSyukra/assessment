package routes

import (
	controller "assesment/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Answer(route *gin.Engine, injector *do.Injector) {
	answerController := do.MustInvoke[controller.AnswerController](injector)

	routes := route.Group("/api/answer")
	{
		routes.POST("", answerController.CreateAnswer)
		routes.GET("", answerController.GetAllAnswers)
		routes.GET("/:id", answerController.GetAnswerByID)
		routes.PUT("/:id", answerController.UpdateAnswer)
		routes.GET("/question/:question_id", answerController.GetAnswerByQuestionID)
		routes.GET("/submission/:submission_id", answerController.GetAnswerBySubmissionID)	
		// routes.GET("/student/:student_id", answerController.GetAnswerByStudentID)
	}
}


