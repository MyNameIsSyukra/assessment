package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	Assessment(server, injector)
	Question(server, injector)
	Answer(server, injector)
}