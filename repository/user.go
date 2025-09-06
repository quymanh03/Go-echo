package repository

import (
	"beginner/db/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (us *UserRepository) FindOne(user *models.User) (*models.User, error) {
	var u models.User
	if err := us.db.Where(user).First(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (us *UserRepository) FindExistedOne(username, email string) (bool, error) {
	var count int64
	err := us.db.Model(&models.User{}).Where("username = ? OR email = ?", username, email).Count(&count).Error
	if count > 0 {
		return true, err
	}
	return false, err
}

func (us *UserRepository) CreateOne(user *models.User) error {
	return us.db.Create(user).Error
}
