package app

import (
	"fmt"
	"github.com/avtara/carehub/utils"
	"github.com/hibiken/asynq"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"log"
)

type App struct {
	Hostname    string
	Server      *echo.Echo
	ZapLogger   *zap.Logger
	DB          *sqlx.DB
	Asynq       *asynq.Client
	AsynqServer *asynq.Server
	AsynqMux    *asynq.ServeMux
	OauthConfig *oauth2.Config
}

func New() App {
	var cfg App

	cfg.Hostname = utils.GetEnv("app.hostname", "127.0.0.1")
	most(cfg.InitViper())
	most(cfg.InitLogger())
	most(cfg.InitAsynq())
	most(cfg.InitServer())
	most(cfg.InitPostgres())
	most(cfg.InitOauth())
	most(cfg.InitService())

	return cfg
}

func (cfg *App) Start() (err error) {
	ch := make(chan bool)
	go func() {
		cfg.Server.Start(fmt.Sprintf("%s:%s", cfg.Hostname, utils.GetEnv("app.port", "8080")))

		ch <- false
	}()

	<-ch
	return
}

func most(err error) {
	if err != nil {
		log.Panic(err)
	}
}
