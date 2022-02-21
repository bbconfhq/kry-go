package handler

import (
	"github.com/labstack/echo/v4"
	"kry-go/model/service"
	"kry-go/server"
	"net/http"
)

type TestcaseHandler struct {
	server  *server.Server
	service *service.TestcaseService
}

func MakeTestcaseHandler(server *server.Server) *TestcaseHandler {
	return &TestcaseHandler{server: server}
}

func (h *TestcaseHandler) GetTestcases(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *TestcaseHandler) PostTestcase(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *TestcaseHandler) GetTestcase(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *TestcaseHandler) PutTestcase(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *TestcaseHandler) DeleteTestcase(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}
