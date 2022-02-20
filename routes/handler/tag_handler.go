package handler

import (
	"github.com/labstack/echo/v4"
	"kry-go/server"
	"kry-go/service"
	"net/http"
)

type TagHandler struct {
	server  *server.Server
	service *service.TagService
}

func MakeTagHandler(server *server.Server) *TagHandler {
	return &TagHandler{server: server}
}

func (tagHandler *TagHandler) GetTag(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}
