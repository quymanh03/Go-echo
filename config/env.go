package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type JwtConfig struct {
	SecretKey  string
	ExpireTime string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Config struct {
	JWT JwtConfig
	DB  DBConfig
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := &Config{
		JWT: JwtConfig{
			SecretKey:  os.Getenv("JWT_SECRET_KEY"),
			ExpireTime: os.Getenv("JWT_EXPIRE_TIME"),
		},
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}
	return config, nil
}
