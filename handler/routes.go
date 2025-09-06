package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(v1 *echo.Group) {
	/* health */
	v1.GET("/health", h.HealthCheck)

	/* auth */
	auth := v1.Group("/auth")
	auth.POST("/login", h.Login)
	auth.POST("/register", h.Register)
}
