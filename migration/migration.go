package migration

import (
	entities "assesment/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error{
	db.AutoMigrate(&entities.Answer{})
	db.AutoMigrate(&entities.Choice{})
	db.AutoMigrate(&entities.Assessment{})
	db.AutoMigrate(&entities.Question{})
	
	return nil
}

func Rollback(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&entities.Answer{}); err != nil {
		return err
	}
	if err := db.Migrator().DropTable(&entities.Choice{}); err != nil {
		return err
	}
	if err := db.Migrator().DropTable(&entities.Assessment{}); err != nil {
		return err
	}
	if err := db.Migrator().DropTable(&entities.Question{}); err != nil {
		return err
	}
	if err := db.Migrator().DropTable(&entities.Submission{}); err != nil {
		return err
	}
	return nil
}