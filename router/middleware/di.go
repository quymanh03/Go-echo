package middleware

import (
	"beginner/config"
	"beginner/utils"

	"github.com/labstack/echo/v4"
)

func DIMiddleware() echo.MiddlewareFunc {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	jwtService := utils.NewJwtService(cfg)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("config", cfg)
			c.Set("jwt", jwtService)
			return next(c)
		}
	}
}
