package utils

import (
	"gorm.io/gorm"
	"kry-go/models"
)

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Submission{})

	if err != nil {
		panic(err.Error())
	}
}
