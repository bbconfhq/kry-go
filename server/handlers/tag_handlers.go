package handlers

import (
	"github.com/labstack/echo/v4"
	"kry-go/server"
	"net/http"
)

type TagHandler struct {
	server *server.Server
}

func MakeTagHandler(server *server.Server) *TagHandler {
	return &TagHandler{server: server}
}

func (tagHandler *TagHandler) GetTag(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}
