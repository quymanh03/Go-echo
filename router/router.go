package router

import (
	"beginner/router/middleware"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()
	e.Pre(middleware.DIMiddleware())
	e.Validator = NewValidator()
	return e
}
