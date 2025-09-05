package handler

import (
	"beginner/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Login
// @Description Login to get JWT token
// @ID auth-login
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body userLoginRequest true "User Login Request"
// @Success 200 {object} utils.HttpResponse
// @Failure 400 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /auth/login [post]
func (h *Handler) Login(c echo.Context) error {
	req := &userLoginRequest{}

	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(err))
	}
	token, err := c.Get("jwt").(*utils.JwtService).GenerateToken(req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err))
	}
	data := newUserLoginResponse(token)
	return c.JSON(http.StatusOK, utils.NewSuccessResponse(data))
}
