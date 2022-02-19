package handlers

import (
	"github.com/labstack/echo/v4"
	"kry-go/models"
	"kry-go/responses"
	"kry-go/server"
	"net/http"
	"strconv"
)

type ContestHandler struct {
	server *server.Server
}

func MakeContestHandler(server *server.Server) *ContestHandler {
	return &ContestHandler{server: server}
}

// GetContests   godoc
// @Summary      Get contests
// @Description  Get contests
// @Tags         contest
// @Accept       json
// @Produce      json
// @Param        page	query	uint	true	"Page of contests"	minimum(1)
// @Success      200	{array}		responses.ContestResponse
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /contest [get]
func (contestHandler *ContestHandler) GetContests(c echo.Context) error {
	var contests []models.Contest
	var contestsResponse []responses.ContestResponse

	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 0)
	if err != nil || page <= 0 {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	offset := int(100 * (page - 1))

	db := contestHandler.server.DB
	db.Preload("Problems").
		Order("ended_at DESC, started_at DESC, id DESC").
		Offset(offset).
		Limit(100).
		Find(&contests)

	for _, e := range contests {
		var problemIds []uint

		for _, f := range e.Problems {
			problemIds = append(problemIds, f.ID)
		}

		contestsResponse = append(
			contestsResponse,
			responses.ContestResponse{
				ID:         e.ID,
				Title:      e.Title,
				Content:    e.Content,
				ProblemIds: problemIds,
				StartedAt:  e.StartedAt,
				EndedAt:    e.EndedAt,
				CreatedAt:  e.CreatedAt,
				UpdatedAt:  e.UpdatedAt,
			},
		)
	}

	return c.JSON(http.StatusOK, contestsResponse)
}

// PostContest   godoc
// @Summary      Post contest
// @Description  Register contest
// @Tags         contest
// @Accept       json
// @Produce      json
// @Param        contest  query  string  true  "string"
// @Success      201	{object}	nil
// @Failure      500	{object}	echo.HTTPError
// @Router       /contest [post]
func (contestHandler *ContestHandler) PostContest(c echo.Context) error {
	db := contestHandler.server.DB
	db.
	return c.JSON(http.StatusCreated, nil)
}

//               param name,param type,data type,is mandatory?,comment attribute(optional)

// GetContest    godoc
// @Summary      Get contest detail
// @Description  Get contest detail
// @Tags         contest
// @Accept       json
// @Produce      json
// @Param        contest_id  path      uint  true  "int valid"
// @Success      200	{object}	models.Contest
// @Failure      500	{int}		http.StatusInternalServerError
// @Router       /contest/:contestId [get]
func (contestHandler *ContestHandler) GetContest(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

// PutContest    godoc
// @Summary      Put contest detail
// @Description  Edit contest detail
// @Tags         contest
// @Accept       json
// @Produce      json
// @Param        contest_id  path   uint  true  "int valid"
// @Success      204         {int}  http.StatusNoContent
// @Failure      404         {int}    http.StatusNotFound
// @Failure      500         {int}    http.StatusInternalServerError
// @Router       /contest/:contestId [put]
func (contestHandler *ContestHandler) PutContest(c echo.Context) error {
	return c.JSON(http.StatusNoContent, nil)
}

// DeleteContest godoc
// @Summary      Delete contest detail
// @Description  Delete contest detail
// @Tags         contest
// @Accept       json
// @Produce      json
// @Param        contest_id  path   uint  true  "int valid"
// @Success      204         {int}  http.StatusNoContent
// @Failure      404         {int}    http.StatusNotFound
// @Failure      500         {int}    http.StatusInternalServerError
// @Router       /contest/:contestId [delete]
func (contestHandler *ContestHandler) DeleteContest(c echo.Context) error {
	return c.JSON(http.StatusNoContent, nil)
}
