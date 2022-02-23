package entity

import (
	"gorm.io/gorm"
	"time"
)

type Contest struct {
	gorm.Model
	Title     string    `gorm:"type:varchar(191);unique;not null"`
	Content   string    `gorm:"type:text;not null"`
	Problems  []Problem `gorm:"many2many:contest_problems;"`
	StartedAt time.Time `gorm:"not null"`
	EndedAt   time.Time `gorm:"not null"`
}
