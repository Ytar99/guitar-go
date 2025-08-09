package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"guitar-go/internal/config"
	"guitar-go/internal/models"
)

type AuthService interface {
	GenerateToken(user *models.User) (string, error)
	ParseToken(tokenString string) (*jwt.Token, error)
	HashPassword(password string) (string, error)
	ComparePassword(hashed, plain string) error
	ValidatePassword(password string) error
}

type authService struct {
	config *config.Config
}

func NewAuthService(cfg *config.Config) AuthService {
	return &authService{config: cfg}
}

func (s *authService) GenerateToken(user *models.User) (string, error) {
	expirationTime, _ := time.ParseDuration(s.config.JWT.Expires)
	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(expirationTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.JWT.Secret))
}

func (s *authService) ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.config.JWT.Secret), nil
	})
}

func (s *authService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *authService) ComparePassword(hashed, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}

func (s *authService) ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}

	return nil
}
