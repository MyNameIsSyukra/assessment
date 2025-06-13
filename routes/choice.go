package routes

import (
	controller "assesment/controller"
	"assesment/middleware"
	"assesment/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Choice(route *gin.Engine, injector *do.Injector) {
	choiceController := do.MustInvoke[controller.ChoiceController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, "jwtService")
	routes := route.Group("/api/v1/choice")
	{
		routes.POST("", middleware.Authenticate(jwtService),choiceController.CreateChoice)
		routes.GET("/",middleware.Authenticate(jwtService), choiceController.GetChoiceByID)
		routes.GET("/question/",middleware.Authenticate(jwtService), choiceController.GetChoicesByQuestionID)
		// routes.PUT("/:id", choiceController.UpdateChoice)
	}
}