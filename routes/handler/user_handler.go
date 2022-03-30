package handler

import (
	"github.com/labstack/echo/v4"
	"kry-go/base"
	"kry-go/model/service"
	"net/http"
)

type UserHandler struct {
	server  *base.Server
	service *service.UserService
}

func MakeUserHandler(server *base.Server) *UserHandler {
	return &UserHandler{server: server}
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *UserHandler) PostUser(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *UserHandler) PutUser(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}
