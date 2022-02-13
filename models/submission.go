package models

import (
	"gorm.io/gorm"
	"time"
)

type Submission struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Result    uint   `gorm:"not null"`
	Language  string `gorm:"not null"`
	Code      string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint
	ProblemID uint
	User      User
}
