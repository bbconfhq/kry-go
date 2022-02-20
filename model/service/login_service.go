package service

import "gorm.io/gorm"

type LoginService struct {
	DB *gorm.DB
}
