package handler

import (
	"beginner/services"

	"gorm.io/gorm"
)

type Handler struct {
	userService services.UserService
}

func New(db *gorm.DB) *Handler {
	return &Handler{
		userService: services.NewUserService(db),
	}
}
