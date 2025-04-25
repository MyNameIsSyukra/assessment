package main

import (
	database "assesment/config"
	middleware "assesment/middleware"
	migration "assesment/migration"
	provider "assesment/provider"
	routes "assesment/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func args(db *gorm.DB) bool {
    if len(os.Args) > 1 {
        if (os.Args[1] == "migrate") {
            migration.Migrate(db)
			print("Migration completed successfully")
            return false
        }
    }
        return true
}


func run(server *gin.Engine) {
	server.Static("/assets", "./assets")


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}

func main() {
    var (
		injector = do.New()
	)
    provider.RegisterDependencies(injector)
    db := database.SetUpDatabaseConnection()
    defer database.CloseDatabaseConnection(db)

    if !args(db) {
		return
	}

    server := gin.Default()
	server.Use(middleware.CORSMiddleware())
    
	// routes
	routes.RegisterRoutes(server, injector)

	run(server)
    
}