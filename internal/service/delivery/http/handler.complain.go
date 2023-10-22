package http

import (
	"github.com/avtara/carehub/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (so *svObject) handlerGetAllCompain(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	limit := int(utils.ToInt(c.QueryParam("limit"), 20))

	response, err := so.complainUseCase.GetAllComplain(ctx, limit)
	if err != nil {
		return err
	}

	return utils.MessageResponse(c, http.StatusOK, response)
}

func (so *svObject) handlerGetComplainByID(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	id := utils.ToInt(c.Param("id"), 0)

	if id <= 0 {
		return utils.MessageResponse(c, http.StatusBadRequest, nil)
	}

	response, err := so.complainUseCase.GetComplainByID(ctx, id)
	if err != nil {
		return err
	}

	return utils.MessageResponse(c, http.StatusOK, response)
}
