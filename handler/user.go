package handler

import (
	"beginner/db/models"
	"beginner/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary GetMe
// @Description Get user information
// @ID user-get-me
// @Tags User
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.HttpResponse
// @Failure 400 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /user/me [get]
func (h *Handler) GetMe(c echo.Context) error {
	userId := c.Get("userId").(string)
	user, err := h.UserRepository.FindOne(&models.User{ID: userId})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err))
	}
	return c.JSON(http.StatusOK, utils.NewSuccessResponse(newGetMeResponse(user)))
}
