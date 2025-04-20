package main

import (
	database "assesment/config"
	migration "assesment/migration"
	"os"

	"gorm.io/gorm"
)

func args(db *gorm.DB) bool {
    if len(os.Args) > 1 {
        if (os.Args[1] == "migrate") {
            print("argadasds")
            migration.Migrate(db)
            return false
        }
    }
        return true
}


func main() {
    db := database.SetUpDatabaseConnection()
    defer database.CloseDatabaseConnection(db)

    if !args(db) {
		return
	}

}