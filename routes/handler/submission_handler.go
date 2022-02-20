package handler

import (
	"github.com/labstack/echo/v4"
	"kry-go/server"
	"kry-go/service"
	"net/http"
)

type SubmissionHandler struct {
	server  *server.Server
	service *service.SubmissionService
}

func MakeSubmissionHandler(server *server.Server) *SubmissionHandler {
	return &SubmissionHandler{server: server}
}

func (submissionHandler *SubmissionHandler) GetSubmission(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}
