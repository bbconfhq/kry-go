package base

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"kry-go/utils"
)

type Server struct {
	Echo *echo.Echo
	DB   *Db
}

func (s *Server) Init(cfg *Config) {
	e := echo.New()
	e.Validator = &utils.Validator{Validator: validator.New()}
	s.Echo = e

	s.DB = &Db{Conn: cfg.Conn}
	s.DB.Init()
}

func (s *Server) Run(address string) error {
	return s.Echo.Start(address)
}
