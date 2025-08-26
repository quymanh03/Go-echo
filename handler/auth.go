package handler

import (
	"beginner/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (*Handler) Login(c echo.Context) error {
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
