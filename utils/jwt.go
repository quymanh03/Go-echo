package utils

import (
	"beginner/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	secretKey  string
	expireTime time.Duration
}

func NewJwtService(cfg *config.Config) *JwtService {
	dur, err := time.ParseDuration(cfg.JWT.ExpireTime)
	if err != nil {
		dur = time.Hour
	}
	return &JwtService{
		secretKey:  cfg.JWT.SecretKey,
		expireTime: dur,
	}
}

func (j *JwtService) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(j.expireTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
}
