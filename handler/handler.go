package handler

import (
	"beginner/repository"

	"gorm.io/gorm"
)

type Handler struct {
	UserRepository repository.UserRepository
}

func New(db *gorm.DB) *Handler {
	return &Handler{
		UserRepository: repository.NewUserRepository(db),
	}
}
