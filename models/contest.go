package models

import (
	"gorm.io/gorm"
	"time"
)

type Contest struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"type:varchar(191);unique;not null"`
	Content   string    `gorm:"type:text;not null"`
	Problems  []Problem `gorm:"many2many:contest_problems;"`
	StartedAt time.Time `gorm:"not null"`
	EndedAt   time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
