package db

import (
	"final-golang-project/models"

	"gorm.io/gorm"
)

func MigrateProductTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.Product{})
}
