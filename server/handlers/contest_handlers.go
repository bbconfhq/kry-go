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

func (contestHandler *ContestHandler) GetContest(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}
