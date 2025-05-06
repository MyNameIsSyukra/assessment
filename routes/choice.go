package routes

import (
	controller "assesment/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Choice(route *gin.Engine, injector *do.Injector) {
	choiceController := do.MustInvoke[controller.ChoiceController](injector)

	routes := route.Group("/api/v1/choice")
	{
		routes.POST("", choiceController.CreateChoice)
		routes.GET("/:id", choiceController.GetChoiceByID)
		routes.GET("/question/:id", choiceController.GetChoicesByQuestionID)
		// routes.PUT("/:id", choiceController.UpdateChoice)
	}
}