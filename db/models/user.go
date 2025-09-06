package models

type User struct {
	ID       string `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Username string `gorm:"column:username;unique;not null"`
	Email    string `gorm:"column:email;unique;not null"`
	Password string `gorm:"column:password;not null"`
}
