package http

import (
	"github.com/avtara/carehub/internal/models"
	"github.com/avtara/carehub/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (so *svObject) handlerEditProfileUser(c echo.Context) error {
	var (
		payload models.UpdateProfileUserParams
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

	err = so.userUseCase.UpdateProfile(ctx, payload, userID)
	if err != nil {
		return err
	}

	return utils.MessageResponse(c, http.StatusOK, nil)
}
