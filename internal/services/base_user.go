package services

import (
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type BaseUserService struct {
	userRepository repositories.UserRepository
}

func NewBaseUserService(userRepository repositories.UserRepository) *BaseUserService {
	return &BaseUserService{
		userRepository: userRepository,
	}
}

func (receiver *BaseUserService) GenerateUUID() string {
	return uuid.New().String()
}

func (receiver *BaseUserService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (receiver *BaseUserService) CheckPasswordByHashAndPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (receiver *BaseUserService) CreateUserByUUIDAndEmailAndPassword(uuid string, email string, password string) error {
	hashedPassword, err := receiver.HashPassword(password)

	if err != nil {
		return err
	}

	user := repositories.NewGormUser()

	user.SetUUID(uuid)
	user.SetEmail(email)
	user.SetPassword(hashedPassword)
	user.SetCreatedAt(int(time.Now().Unix()))
	user.SetUpdatedAt(int(time.Now().Unix()))

	err = receiver.userRepository.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}
