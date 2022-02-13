package handlers

import (
	"github.com/labstack/echo/v4"
	"kry-go/server"
	"net/http"
)

type TestcaseHandler struct {
	server *server.Server
}

func MakeTestcaseHandler(server *server.Server) *TestcaseHandler {
	return &TestcaseHandler{server: server}
}

func (testcaseHandler *TestcaseHandler) GetTestcase(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}
