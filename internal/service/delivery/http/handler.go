package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (so *svObject) handlerHelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"data": "hello, world!",
	})
}
