package handler

import (
	"github.com/labstack/echo/v4"
	"kry-go/model/service"
	"kry-go/server"
	"net/http"
)

type TagHandler struct {
	server  *server.Server
	service *service.TagService
}

func MakeTagHandler(server *server.Server) *TagHandler {
	return &TagHandler{server: server}
}

func (h *TagHandler) GetTags(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *TagHandler) PostTag(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *TagHandler) GetTag(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *TagHandler) PutTag(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *TagHandler) DeleteTag(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}
