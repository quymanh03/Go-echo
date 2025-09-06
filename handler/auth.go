package handler

import (
	"beginner/db/models"
	"beginner/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	user, err := h.UserRepository.FindOne(&models.User{Email: req.Email})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err))
	}

	if user == nil {
		return c.JSON(http.StatusBadRequest, utils.NewCustomErrorResponse("Invalid email or password"))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewCustomErrorResponse("Invalid email or password"))
	}

	token, err := c.Get("jwt").(*utils.JwtService).GenerateToken(user.ID.String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err))
	}
	data := newUserLoginResponse(token)
	return c.JSON(http.StatusOK, utils.NewSuccessResponse(data))
}

// @Summary Register
// @Description Register a new user
// @ID auth-register
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body userRegisterRequest true "User Register Request"
// @Success 200 {object} utils.HttpResponse
// @Failure 400 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /auth/register [post]
func (h *Handler) Register(c echo.Context) error {
	req := &userRegisterRequest{}

	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(err))
	}

	isExisted, err := h.UserRepository.FindExistedOne(req.Username, req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err))
	}

	if isExisted {
		return c.JSON(http.StatusBadRequest, utils.NewCustomErrorResponse("Email or usename has already existed!"))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err))
	}

	err = h.UserRepository.CreateOne(&models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err))
	}

	return c.JSON(http.StatusCreated, utils.NewSuccessResponse("Register successfully!"))
}
