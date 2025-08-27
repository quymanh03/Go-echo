package router

import (
	"beginner/router/middleware"

	_ "beginner/docs"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	e.Pre(middleware.DIMiddleware())
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
	}))
	e.Validator = NewValidator()
	return e
}
