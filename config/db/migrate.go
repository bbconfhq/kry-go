package db

import (
	"gorm.io/gorm"
	"kry-go/entity"
)

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.Contest{},
		&entity.Problem{},
		&entity.Submission{},
		&entity.Tag{},
		&entity.Testcase{},
		&entity.User{})

	if err != nil {
		panic(err.Error())
	}
}
