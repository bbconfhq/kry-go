package handler

import (
	"github.com/labstack/echo/v4"
	"kry-go/model/request"
	"kry-go/model/service"
	"kry-go/server"
	"net/http"
	"strconv"
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
// @Param         page		query	uint						true	"Page of submissions"	minimum(1)
// @Param         user_id	query	uint						false	"User ID"
// @Param         language	query	string						false	"Type of language"
// @Success       200	{array}		response.SubmissionResponse
// @Failure       400	{object}	echo.HTTPError
// @Failure       500	{object}	echo.HTTPError
// @Router        /submission [get]
func (h *SubmissionHandler) GetSubmissions(c echo.Context) error {
	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 0)

	if err != nil || page <= 0 {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := strconv.ParseInt(c.QueryParam("user_id"), 10, 0)
	if err != nil {
		id = -1
	}

	language := c.QueryParam("language")

	offset := int(100 * (page - 1))
	result, err := h.service.SelectSubmissionsByFilter(offset, int(id), language, 100)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
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
	var submissionRequest request.SubmissionRequest

	if err := (&echo.DefaultBinder{}).BindBody(c, &submissionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.service.CreateSubmission(&submissionRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

// GetSubmission godoc
// @Summary      Get submission detail
// @Description  Get submission detail
// @Tags         Submission Endpoints
// @Accept       json
// @Produce      json
// @Param        id			path	uint						true	"Submission ID"
// @Success      200	{object}	response.SubmissionResponse
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /submission/{id} [get]
func (h *SubmissionHandler) GetSubmission(c echo.Context) error {
	id, err := strconv.ParseInt(c.QueryParam("id"), 10, 0)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.service.SelectSubmission(int(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
