package utils

import (
	"gorm.io/gorm"
	"kry-go/models"
)

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Contest{},
		&models.Problem{},
		&models.Submission{},
		&models.Tag{},
		&models.Testcase{},
		&models.User{})

	if err != nil {
		panic(err.Error())
	}
}
