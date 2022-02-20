package service

import "gorm.io/gorm"

type ProblemService struct {
	DB *gorm.DB
}
