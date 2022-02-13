package handlers

import (
	"github.com/labstack/echo/v4"
	"kry-go/server"
	"net/http"
)

type SubmissionHandler struct {
	server *server.Server
}

func MakeSubmissionHandler(server *server.Server) *SubmissionHandler {
	return &SubmissionHandler{server: server}
}

func (submissionHandler *SubmissionHandler) GetSubmission(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}
