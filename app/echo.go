package app

import (
	"github.com/avtara/carehub/pkg"
	"github.com/avtara/carehub/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (cfg *App) InitServer() (err error) {
	cfg.Server = echo.New()
	cfg.Server.Use(
		pkg.ZapLogger(cfg.ZapLogger),
		middleware.Recover(),
		middleware.RequestID(),
		session.Middleware(
			sessions.NewCookieStore(
				[]byte(utils.GetEnv("session.id", "!s3cret!")),
			),
		),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
		}),
	)
	cfg.Server.Validator = &pkg.CustomValidator{Validator: validator.New()}
	cfg.Server.HTTPErrorHandler = utils.CustomHTTPErrorHandler

	return
}
