package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Username string    `gorm:"column:username;unique;not null"`
	Email    string    `gorm:"column:email;unique;not null"`
	Password string    `gorm:"column:password;not null"`
}
