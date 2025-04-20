package migration

import (
	entities "assesment/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error{
	db.AutoMigrate(&entities.Answer{})
	db.AutoMigrate(&entities.Choice{})
	db.AutoMigrate(&entities.Evaluation{})
	db.AutoMigrate(&entities.Question{})
	
	return nil
}