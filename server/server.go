package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kry-go/config"
	"kry-go/config/db"
	"kry-go/utils"
)

type Server struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func Init(cfg *config.Config) *Server {
	e := echo.New()
	e.Validator = &utils.Validator{Validator: validator.New()}
	return &Server{
		Echo: e,
		DB:   db.Init(cfg),
	}
}

func (server *Server) Run(address string) error {
	return server.Echo.Start(address)
}
