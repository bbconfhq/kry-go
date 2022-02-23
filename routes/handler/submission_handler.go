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

// GetSubmissions godoc
// @Summary       Get submissions
// @Description   Get submissions
// @Tags          Submission Endpoints
// @Accept        json
// @Produce       json
// @Param         page		query	uint	true	"Page of submissions"	minimum(1)
// @Param         user_id	query	string	false	"User ID"
// @Param         language	query	string	false	"Type of language"
// @Success       200	{array}		response.SubmissionResponse
// @Failure       400	{object}	echo.HTTPError
// @Failure       500	{object}	echo.HTTPError
// @Router        /submission [get]
func (h *SubmissionHandler) GetSubmissions(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

// PostSubmission godoc
// @Summary       Post submission
// @Description   Register submission
// @Tags          Submission Endpoints
// @Accept        json
// @Produce       json
// @Param         contest	body	request.SubmissionRequest	true	"Necessary submission information"
// @Success       201	{object}	nil
// @Failure       400	{object}	echo.HTTPError
// @Failure       500	{object}	echo.HTTPError
// @Router        /submission [post]
func (h *SubmissionHandler) PostSubmission(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

// GetSubmission godoc
// @Summary      Get submission detail
// @Description  Get submission detail
// @Tags         Submission Endpoints
// @Accept       json
// @Produce      json
// @Param        submission_id	path	uint					true	"Submission ID"
// @Success      200	{object}	response.SubmissionResponse
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /submission/{submission_id} [get]
func (h *SubmissionHandler) GetSubmission(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}
