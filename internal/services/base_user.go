package services

import (
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type BaseUserService struct {
	userRepository repositories.UserRepository
	userBuilder    repositories.UserBuilder
}

func NewBaseUserService(
	userRepository repositories.UserRepository,
	userBuilder repositories.UserBuilder,
) *BaseUserService {
	return &BaseUserService{
		userRepository: userRepository,
		userBuilder:    userBuilder,
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

func (receiver *BaseUserService) CheckHashedPasswordAndNativePassword(
	hashedPassword string,
	nativePassword string,
) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(nativePassword))
}

func (receiver *BaseUserService) CreateUserByUUIDAndEmailAndHashedPassword(
	uuid string,
	email string,
	hashedPassword string,
) error {
	newUser, err := receiver.userBuilder.
		New().
		SetUUID(uuid).
		SetEmail(email).
		SetPassword(hashedPassword).
		SetCreatedAt(int(time.Now().Unix())).
		SetUpdatedAt(int(time.Now().Unix())).
		Build()

	if err != nil {
		return err
	}

	err = receiver.userRepository.CreateUser(newUser)

	if err != nil {
		return err
	}

	return nil
}

func (receiver *BaseUserService) UpdatePasswordByUUIDAndHashedPassword(
	uuid string,
	hashedPassword string,
) error {
	err := receiver.userRepository.UpdatePasswordByUUIDAndPasswordAndUpdatedAt(
		uuid,
		hashedPassword,
		int(time.Now().Unix()),
	)

	if err != nil {
		return err
	}

	return nil
}
