package services

import (
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (receiver *MockUserService) GenerateUUID() string {
	return receiver.Called().String(0)
}

func (receiver *MockUserService) HashPassword(password string) (string, error) {
	args := receiver.Called(password)

	return args.String(0), args.Error(1)
}

func (receiver *MockUserService) CheckPasswordByHashAndPassword(hashedPassword string, password string) error {
	return receiver.Called(hashedPassword, password).Error(0)
}

type MockResetTokenRepository struct {
	mock.Mock
}

func (receiver *MockResetTokenRepository) HasToken(token string) bool {
	return receiver.Called(token).Bool(0)
}

func (receiver *MockResetTokenRepository) GetResetTokenByToken(token string) repositories.ResetToken {
	args := receiver.Called(token)

	if args.Get(0) != nil {
		return args.Get(0).(repositories.ResetToken)
	}

	return nil
}

func (receiver *MockResetTokenRepository) CreateResetToken(token repositories.ResetToken) error {
	return receiver.Called(token).Error(0)
}

func (receiver *MockResetTokenRepository) DeleteResetToken(token string) error {
	return receiver.Called(token).Error(0)
}

type MockUserRepository struct {
	mock.Mock
}

func (receiver *MockUserRepository) HasUserByUUID(uuid string) bool {
	return receiver.Called(uuid).Bool(0)
}

func (receiver *MockUserRepository) HasUserByEmail(email string) bool {
	return receiver.Called(email).Bool(0)
}

func (receiver *MockUserRepository) HasUserByEmailAndPassword(email string, password string) bool {
	return receiver.Called(email, password).Bool(0)
}

func (receiver *MockUserRepository) HasUserByUUIDAndPassword(uuid string, password string) bool {
	return receiver.Called(uuid, password).Bool(0)
}

func (receiver *MockUserRepository) GetUserByEmail(email string) repositories.User {
	args := receiver.Called(email)

	if args.Get(0) != nil {
		return args.Get(0).(repositories.User)
	}

	return nil
}

func (receiver *MockUserRepository) CreateUser(user repositories.User) error {
	return receiver.Called(user).Error(0)
}

func (receiver *MockUserRepository) UpdatePassword(uuid string, password string, updatedAt int) error {
	return receiver.Called(uuid, password, updatedAt).Error(0)
}
