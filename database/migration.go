package database

import (
	"go-agreenery/models"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.Credential{},
		&models.User{},
		&models.Category{},
		&models.Plant{},
		&models.Step{},
		&models.EnrolledPlant{},
		&models.EnrolledStep{},
		&models.Post{},
		&models.Comment{},
		&models.Like{},
		&models.Article{},
		&models.PostReport{},
	)
}
