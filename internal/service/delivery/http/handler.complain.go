package http

import (
	"github.com/avtara/carehub/internal/models"
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

func (so *svObject) handlerInsertComplain(c echo.Context) error {
	var (
		payload models.InsertComplainParams
		ctx     = c.Request().Context()
	)

	data, err := utils.ExtractTokenJWT(c)
	if err != nil {
		return err
	}

	userID := utils.ToInt(data["user_id"], 0)

	if err = c.Bind(&payload); err != nil {
		return utils.MessageResponse(c, http.StatusBadRequest, nil)
	}

	if err = c.Validate(payload); err != nil {
		return utils.MessageResponse(c, http.StatusBadRequest, nil)
	}

	err = so.complainUseCase.InsertComplain(ctx, payload, userID)
	if err != nil {
		return err
	}

	return utils.MessageResponse(c, http.StatusOK, nil)
}

func (so *svObject) handlerInsertResolution(c echo.Context) error {
	var (
		payload models.UpdateStatusComplainParams
		ctx     = c.Request().Context()
	)

	complainID := utils.ToInt(c.Param("id"), 0)

	if err := c.Bind(&payload); err != nil {
		return utils.MessageResponse(c, http.StatusBadRequest, nil)
	}

	if err := c.Validate(payload); err != nil {
		return utils.MessageResponse(c, http.StatusBadRequest, nil)
	}

	err := so.complainUseCase.UpdateStatus(ctx, payload.Status, complainID)
	if err != nil {
		return err
	}

	return utils.MessageResponse(c, http.StatusOK, nil)
}
