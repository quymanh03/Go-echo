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

type Config struct {
	JWT JwtConfig
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
	}
	return config, nil
}
