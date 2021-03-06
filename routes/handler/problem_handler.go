package handler

import (
	"github.com/labstack/echo/v4"
	"kry-go/base"
	"kry-go/model/request"
	"kry-go/model/service"
	"net/http"
	"strconv"
)

type ProblemHandler struct {
	server  *base.Server
	service *service.ProblemService
}

func MakeProblemHandler(server *base.Server) *ProblemHandler {
	return &ProblemHandler{
		server:  server,
		service: &service.ProblemService{DB: server.DB.Db},
	}
}

// GetProblems   godoc
// @Summary      Get problems
// @Description  Get problems
// @Tags         Problem Endpoints
// @Accept       json
// @Produce      json
// @Param        page		query	uint					true	"Page of problems"	minimum(1)
// @Success      200	{array}		response.ProblemResponse
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /problem [get]
func (h *ProblemHandler) GetProblems(c echo.Context) error {
	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 0)
	if err != nil || page <= 0 {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	offset := int(100 * (page - 1))
	result, err := h.service.SelectProblemsAt(offset, 100)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// PostProblem   godoc
// @Summary      Post problem
// @Description  Register problem
// @Tags         Problem Endpoints
// @Accept       json
// @Produce      json
// @Param        contest	body	request.ProblemRequest	true	"Necessary problem information"
// @Success      201	{object}	nil
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /problem [post]
func (h *ProblemHandler) PostProblem(c echo.Context) error {
	var problemRequest request.ProblemRequest

	if err := (&echo.DefaultBinder{}).BindBody(c, &problemRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.service.CreateProblem(&problemRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

// GetProblem    godoc
// @Summary      Get problem detail
// @Description  Get problem detail
// @Tags         Problem Endpoints
// @Accept       json
// @Produce      json
// @Param        id			path	uint					true	"Problem ID"
// @Success      200	{object}	response.ProblemResponse
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /problem/{problem_id} [get]
func (h *ProblemHandler) GetProblem(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.service.SelectProblem(int(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// PutProblem    godoc
// @Summary      Put problem detail
// @Description  Edit problem detail
// @Tags         Problem Endpoints
// @Accept       json
// @Produce      json
// @Param        problem	body	request.ContestRequest	true	"Necessary problem information"
// @Success      201	{object}	nil
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /problem/{id} [put]
func (h *ProblemHandler) PutProblem(c echo.Context) error {
	var problemRequest request.ProblemRequest

	if err := (&echo.DefaultBinder{}).BindBody(c, &problemRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.service.UpdateProblem(&problemRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

// DeleteProblem godoc
// @Summary      Delete problem detail
// @Description  Delete problem detail
// @Tags         Problem Endpoints
// @Accept       json
// @Produce      json
// @Param        id			path	uint					true	"Problem ID"
// @Success      204	{object}	nil
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /problem/{id} [delete]
func (h *ProblemHandler) DeleteProblem(c echo.Context) error {
	id, err := strconv.ParseInt(c.QueryParam("id"), 10, 0)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.service.DeleteProblem(int(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
