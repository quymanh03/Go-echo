package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase(cfg *DBConfig) error {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}

func DB() *gorm.DB {
	return db
}
