package handler

import (
	"github.com/labstack/echo/v4"
	"kry-go/model/request"
	"kry-go/model/service"
	"kry-go/server"
	"net/http"
	"strconv"
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
	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 0)

	if err != nil || page <= 0 {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	offset := int(100 * (page - 1))
	result, err := h.service.SelectTagsAt(offset, 100)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
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
	var tagRequest request.TagRequest

	if err := (&echo.DefaultBinder{}).BindBody(c, &tagRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.service.CreateTag(&tagRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
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
	id, err := strconv.ParseInt(c.QueryParam("id"), 10, 0)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.service.SelectTag(int(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
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
	return c.JSON(http.StatusCreated, nil)
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
	return c.JSON(http.StatusNoContent, nil)
}
