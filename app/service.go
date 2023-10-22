package app

import (
	broker2 "github.com/avtara/carehub/internal/service/delivery/broker"
	"github.com/avtara/carehub/internal/service/delivery/http"
	"github.com/avtara/carehub/internal/service/repository/broker"
	"github.com/avtara/carehub/internal/service/repository/postgres"
	"github.com/avtara/carehub/internal/service/usecase"
)

func (cfg *App) InitService() (err error) {
	userRepository := postgres.NewUserRepository(cfg.DB)
	brokerRepository := broker.NewBrokerRepository(cfg.Asynq)
	userUseCase := usecase.NewUserUseCase(userRepository, brokerRepository, cfg.DB)
	authenticationUseCase := usecase.NewAuthenticationUseCase(userRepository, brokerRepository, cfg.OauthConfig, cfg.DB)

	http.NewServerHandler(cfg.Server, cfg.ZapLogger, authenticationUseCase, userUseCase)
	broker2.NewBrokerHandler(cfg.Asynq, cfg.AsynqServer, cfg.AsynqMux)

	return
}
