package handler

import (
	"github.com/labstack/echo/v4"
	"kry-go/model/service"
	"kry-go/server"
	"net/http"
)

type SubmissionHandler struct {
	server  *server.Server
	service *service.SubmissionService
}

func MakeSubmissionHandler(server *server.Server) *SubmissionHandler {
	return &SubmissionHandler{server: server}
}

func (h *SubmissionHandler) GetSubmissions(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *SubmissionHandler) PostSubmission(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *SubmissionHandler) GetSubmission(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *SubmissionHandler) PutSubmission(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

func (h *SubmissionHandler) DeleteSubmission(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}
