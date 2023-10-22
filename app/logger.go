package app

import (
	"github.com/avtara/carehub/utils"
	"go.uber.org/zap"
)

func (cfg *App) InitLogger() (err error) {
	switch utils.GetEnv("app.env", "development") {
	case "production":
		logConfig := zap.NewProductionConfig()
		logConfig.DisableStacktrace = true
		logConfig.DisableCaller = true
		cfg.ZapLogger, err = logConfig.Build()
	default:
		logConfig := zap.NewDevelopmentConfig()
		logConfig.DisableStacktrace = true
		logConfig.DisableCaller = true
		cfg.ZapLogger, err = logConfig.Build()
	}
	defer cfg.ZapLogger.Sync()

	return
}
