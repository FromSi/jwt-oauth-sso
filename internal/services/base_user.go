package services

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type BaseUserService struct{}

func NewBaseUserService() *BaseUserService {
	return &BaseUserService{}
}

func (receiver BaseUserService) GenerateUUID() string {
	return uuid.New().String()
}

func (receiver BaseUserService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (receiver BaseUserService) CheckPasswordByHashAndPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
