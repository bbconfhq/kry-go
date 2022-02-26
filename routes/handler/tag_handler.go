package handler

import (
	"github.com/labstack/echo/v4"
	"kry-go/model/service"
	"kry-go/server"
	"net/http"
)

type TagHandler struct {
	server  *server.Server
	service *service.TagService
}

func MakeTagHandler(server *server.Server) *TagHandler {
	return &TagHandler{server: server}
}

// GetTags godoc
// @Summary       Get tags
// @Description   Get tags
// @Tags          Tag Endpoints
// @Accept        json
// @Produce       json
// @Param         page		query	uint	true	"Page of tags"	minimum(1)
// @Success       200	{array}		response.TagResponse
// @Failure       400	{object}	echo.HTTPError
// @Failure       500	{object}	echo.HTTPError
// @Router        /tag [get]
func (h *TagHandler) GetTags(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

// PostTag godoc
// @Summary       Post tag
// @Description   Register tag
// @Tags          Tag Endpoints
// @Accept        json
// @Produce       json
// @Param         tag	body	request.TagRequest	true	"Necessary tag information"
// @Success       201	{object}	nil
// @Failure       400	{object}	echo.HTTPError
// @Failure       500	{object}	echo.HTTPError
// @Router        /tag [post]
func (h *TagHandler) PostTag(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

// GetTag godoc
// @Summary      Get tag detail
// @Description  Get tag detail
// @Tags         Tag Endpoints
// @Accept       json
// @Produce      json
// @Param        id	path	uint					true	"Tag ID"
// @Success      200	{object}	response.TagResponse
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /tag/{id} [get]
func (h *TagHandler) GetTag(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

// PutTag    godoc
// @Summary      Put tag detail
// @Description  Edit tag detail
// @Tags         Tag Endpoints
// @Accept       json
// @Produce      json
// @Param        problem	body	request.TagRequest	true	"Necessary tag information"
// @Success      201	{object}	nil
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /tag/{id} [put]
func (h *TagHandler) PutTag(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

// DeleteTag godoc
// @Summary      Delete tag detail
// @Description  Delete tag detail
// @Tags         Tag Endpoints
// @Accept       json
// @Produce      json
// @Param        id	path	uint					true	"Tag ID"
// @Success      204	{object}	nil
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /tag/{id} [delete]
func (h *TagHandler) DeleteTag(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}
