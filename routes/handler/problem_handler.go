package handler

import (
	"github.com/labstack/echo/v4"
	"kry-go/server"
	"kry-go/service"
	"net/http"
)

type ProblemHandler struct {
	server  *server.Server
	service *service.ProblemService
}

func MakeProblemHandler(server *server.Server) *ProblemHandler {
	return &ProblemHandler{server: server}
}

func (problemHandler *ProblemHandler) GetProblem(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}
