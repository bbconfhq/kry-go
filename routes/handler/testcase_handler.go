package handler

import (
	"github.com/labstack/echo/v4"
	"kry-go/model/service"
	"kry-go/server"
	"net/http"
	"strconv"
)

type TestcaseHandler struct {
	server  *server.Server
	service *service.TestcaseService
}

func MakeTestcaseHandler(server *server.Server) *TestcaseHandler {
	return &TestcaseHandler{server: server}
}

// GetTestcases   godoc
// @Summary       Get testcases
// @Description   Get testcases
// @Tags          Testcase Endpoints
// @Accept        json
// @Produce       json
// @Param         problem_id	query	uint					true	"Problem ID"	minimum(1)
// @Success       200	{array}		response.TestcaseResponse
// @Failure       400	{object}	echo.HTTPError
// @Failure       500	{object}	echo.HTTPError
// @Router        /testcase [get]
func (h *TestcaseHandler) GetTestcases(c echo.Context) error {
	id, err := strconv.ParseInt(c.QueryParam("problem_id"), 10, 0)

	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.service.SelectTestcasesAt(int(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// PostTestcase   godoc
// @Summary       Post testcase
// @Description   Register testcase
// @Tags          Testcase Endpoints
// @Accept        json
// @Produce       json
// @Param         testcase	body	request.TestcaseRequest	true	"Necessary testcase information"
// @Success       201	{object}	nil
// @Failure       400	{object}	echo.HTTPError
// @Failure       500	{object}	echo.HTTPError
// @Router        /testcase [post]
func (h *TestcaseHandler) PostTestcase(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

// GetTestcase   godoc
// @Summary      Get testcase detail
// @Description  Get testcase detail
// @Tags         Testcase Endpoints
// @Accept       json
// @Produce      json
// @Param        id			path	uint					true	"Testcase ID"
// @Success      200	{object}	response.TestcaseResponse
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /testcase/{id} [get]
func (h *TestcaseHandler) GetTestcase(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

// PutTestcase   godoc
// @Summary      Put testcase detail
// @Description  Edit testcase detail
// @Tags         Testcase Endpoints
// @Accept       json
// @Produce      json
// @Param        testcase	body	request.TestcaseRequest	true	"Necessary testcase information"
// @Success      201	{object}	nil
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /testcase/{id} [put]
func (h *TestcaseHandler) PutTestcase(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}

// DeleteTestcase godoc
// @Summary      Delete testcase detail
// @Description  Delete testcase detail
// @Tags         Testcase Endpoints
// @Accept       json
// @Produce      json
// @Param        id			path	uint					true	"Testcase ID"
// @Success      204	{object}	nil
// @Failure      400	{object}	echo.HTTPError
// @Failure      500	{object}	echo.HTTPError
// @Router       /testcase/{id} [delete]
func (h *TestcaseHandler) DeleteTestcase(c echo.Context) error {
	return c.JSON(http.StatusOK, 0)
}
