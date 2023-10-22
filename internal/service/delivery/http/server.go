package http

import (
	"github.com/avtara/carehub/internal/service"
	"github.com/avtara/carehub/utils"
	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type svObject struct {
	service               *echo.Echo
	asynq                 *asynq.Client
	logger                *zap.Logger
	authenticationUseCase service.AuthenticationUseCase
	userUseCase           service.UserUseCase
	categoryUseCase       service.CategoryUseCase

	IsSystemMaintenance bool
}

func NewServerHandler(
	svc *echo.Echo,
	logger *zap.Logger,
	authenticationUsecase service.AuthenticationUseCase,
	userUsecase service.UserUseCase,
	categoryUseCase service.CategoryUseCase,
) {
	obj := &svObject{
		service:               svc,
		logger:                logger,
		authenticationUseCase: authenticationUsecase,
		userUseCase:           userUsecase,
		categoryUseCase:       categoryUseCase,

		IsSystemMaintenance: utils.ToBool(utils.GetEnv("app.is_system_maintenance", "FALSE"), false) == true,
	}

	obj.initRoute()
}
