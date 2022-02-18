package handlers

import (
	"github.com/labstack/echo/v4"
	"kry-go/server"
	"net/http"
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
// @Success      200  {array}	models.Contest
// @Failure      500  {int}  	http.StatusInternalServerError
// @Router       /contest [get]
func (contestHandler *ContestHandler) GetContests(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

// PostContest   godoc
// @Summary      Post contest
// @Description  Register contest
// @Tags         contest
// @Accept       json
// @Produce      json
// @Param        contest	query	string	true	"string"
// @Success      204  {int}     http.StatusNoContent
// @Failure      500  {int}  	http.StatusInternalServerError
// @Router       /contest [post]
func (contestHandler *ContestHandler) PostContest(c echo.Context) error {
	return c.JSON(http.StatusNoContent, nil)
}

//               param name,param type,data type,is mandatory?,comment attribute(optional)

// GetContest    godoc
// @Summary      Get contest detail
// @Description  Get contest detail
// @Tags         contest
// @Accept       json
// @Produce      json
// @Param        contest_id	path	uint	true	"int valid"
// @Success      200  {object}  models.Contest
// @Failure      500  {int}  	http.StatusInternalServerError
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
// @Param        contest_id	path	uint	true	"int valid"
// @Success      204  {int}     http.StatusNoContent
// @Failure      404  {int}  	http.StatusNotFound
// @Failure      500  {int}  	http.StatusInternalServerError
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
// @Param        contest_id	path	uint	true	"int valid"
// @Success      204  {int}     http.StatusNoContent
// @Failure      404  {int}  	http.StatusNotFound
// @Failure      500  {int}  	http.StatusInternalServerError
// @Router       /contest/:contestId [delete]
func (contestHandler *ContestHandler) DeleteContest(c echo.Context) error {
	return c.JSON(http.StatusNoContent, nil)
}
