package middleware

import (
	"beginner/utils"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
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
			decoded, err := jwtService.ValidateToken(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, utils.NewErrorResponse(err))
			}

			if _, ok := decoded.Claims.(jwt.MapClaims); !ok {
				return c.JSON(http.StatusUnauthorized, utils.NewCustomErrorResponse("Invalid token claims"))
			}
			c.Set("userId", decoded.Claims.(jwt.MapClaims)["user_id"])
			return next(c)
		}
	}
}

func extractTokenFromHeader(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		return authHeader[7:], nil
	}
	return "", echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
}
