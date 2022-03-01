package handler

import (
	"github.com/labstack/echo/v4"
	request "kry-go/model/request"
	"kry-go/model/service"
	"kry-go/server"
	"net/http"
	"strconv"
)

type ContestHandler struct {
	server  *server.Server
	service *service.ContestService
}

func MakeContestHandler(server *server.Server) *ContestHandler {
	return &ContestHandler{
		server:  server,
		service: &service.ContestService{DB: server.DB},
	}
}

// GetContests   godoc
// @Summary      Get contests
// @Description  Get contests
// @Tags         Contest Endpoints
// @Accept       json
// @Produce      json
// @Param        page		query	uint					true	"Page of contests"	minimum(1)
// @Success      200	{array}		response.ContestResponse
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /contest [get]
func (h *ContestHandler) GetContests(c echo.Context) error {
	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 0)
	if err != nil || page <= 0 {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	offset := int(100 * (page - 1))
	result, err := h.service.SelectContestsAt(offset, 100)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// PostContest   godoc
// @Summary      Post contest
// @Description  Register contest
// @Tags         Contest Endpoints
// @Accept       json
// @Produce      json
// @Param        contest	body	request.ContestRequest	true	"Necessary contest information"
// @Success      201	{object}	nil
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /contest [post]
func (h *ContestHandler) PostContest(c echo.Context) error {
	var contestRequest request.ContestRequest

	if err := (&echo.DefaultBinder{}).BindBody(c, &contestRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.service.CreateContest(&contestRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

// GetContest    godoc
// @Summary      Get contest detail
// @Description  Get contest detail
// @Tags         Contest Endpoints
// @Accept       json
// @Produce      json
// @Param        id			path	uint					true	"Contest ID"
// @Success      200	{object}	response.ContestResponse
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /contest/{id} [get]
func (h *ContestHandler) GetContest(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.service.SelectContest(int(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// PutContest    godoc
// @Summary      Put contest detail
// @Description  Edit contest detail
// @Tags         Contest Endpoints
// @Accept       json
// @Produce      json
// @Param        contest	body	request.ContestRequest	true	"Necessary contest information"
// @Success      201	{object}	nil
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /contest/{id} [put]
func (h *ContestHandler) PutContest(c echo.Context) error {
	var contestRequest request.ContestRequest

	if err := (&echo.DefaultBinder{}).BindBody(c, &contestRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.service.UpdateContest(&contestRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

// DeleteContest godoc
// @Summary      Delete contest detail
// @Description  Delete contest detail
// @Tags         Contest Endpoints
// @Accept       json
// @Produce      json
// @Param        id			path	uint					true	"Contest ID"
// @Success      204	{object}	nil
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /contest/{id} [delete]
func (h *ContestHandler) DeleteContest(c echo.Context) error {
	id, err := strconv.ParseInt(c.QueryParam("id"), 10, 0)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.service.DeleteContest(int(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
