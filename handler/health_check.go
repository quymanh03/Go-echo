package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Health Check
// @Description Health Check
// @ID get-health
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.HttpResponse
// @Failure 400 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /health [get]
func (h *Handler) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
