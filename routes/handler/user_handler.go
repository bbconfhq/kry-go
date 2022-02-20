package handler

import (
	"github.com/labstack/echo/v4"
	"kry-go/model/service"
	"kry-go/server"
	"net/http"
)

type UserHandler struct {
	server  *server.Server
	service *service.UserService
}

func MakeUserHandler(server *server.Server) *UserHandler {
	return &UserHandler{server: server}
}

func (userHandler *UserHandler) GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}
