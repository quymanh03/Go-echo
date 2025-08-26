package middleware

import (
	"beginner/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func JwtMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := extractTokenFromHeader(c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, utils.NewErrorResponse(err))
			}
			jwtService := c.Get("jwt").(*utils.JwtService)
			_, err = jwtService.ValidateToken(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, utils.NewErrorResponse(err))
			}
			return next(c)
		}
	}
}

func extractTokenFromHeader(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	log.Println("Auth Header:", authHeader[:7])
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		return authHeader[7:], nil
	}
	return "", echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
}
