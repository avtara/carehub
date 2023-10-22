package http

import (
	"github.com/avtara/carehub/internal/models"
	"github.com/avtara/carehub/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (so *svObject) handlerAuthenticationGoogle(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	response, err := so.authenticationUseCase.AuthenticationGoogle(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"redirect_url": response,
	})
}

func (so *svObject) handlerCallbackGoogle(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	code := c.QueryParam("code")

	response, err := so.authenticationUseCase.GoogleCallback(ctx, code)
	if err != nil {
		return err
	}

	return utils.MessageResponse(c, http.StatusOK, response)
}

func (so *svObject) handlerAuthenticationLogin(c echo.Context) error {
	var (
		payload models.AuthenticationParams
		ctx     = c.Request().Context()
	)

	if err := c.Bind(&payload); err != nil {
		return utils.MessageResponse(c, http.StatusBadRequest, nil)
	}

	if err := c.Validate(payload); err != nil {
		return utils.MessageResponse(c, http.StatusBadRequest, nil)
	}

	response, err := so.authenticationUseCase.Login(ctx, payload)
	if err != nil {
		return err
	}

	return utils.MessageResponse(c, http.StatusOK, response)
}
