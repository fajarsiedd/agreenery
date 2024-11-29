package database

import (
	"go-agreenery/models"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.Credential{}, &models.User{})

	// Add table suffix when creating tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Credential{}, &models.User{})
}
