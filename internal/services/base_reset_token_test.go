package services

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
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

func (receiver *MockUserRepository) GetUserByEmailAndPassword(email string, password string) repositories.User {
	args := receiver.Called(email, password)

	if args.Get(0) != nil {
		return args.Get(0).(repositories.User)
	}

	return nil
}

func (receiver *MockUserRepository) GetUserByUUIDAndPassword(uuid string, password string) repositories.User {
	args := receiver.Called(uuid, password)

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

func Test_NewBaseResetTokenService(t *testing.T) {
	mockUserService := MockUserService{}
	mockUserRepository := MockUserRepository{}
	mockResetTokenRepository := MockResetTokenRepository{}

	baseResetTokenService := NewBaseResetTokenService(&mockUserService, &mockResetTokenRepository, &mockUserRepository)

	assert.NotNil(t, baseResetTokenService)
}

func TestBaseResetTokenService_GenerateToken(t *testing.T) {
	mockUserService := MockUserService{}
	mockUserRepository := MockUserRepository{}
	mockResetTokenRepository := MockResetTokenRepository{}

	baseResetTokenService := NewBaseResetTokenService(&mockUserService, &mockResetTokenRepository, &mockUserRepository)

	uuidOne := baseResetTokenService.GenerateToken()
	uuidTwo := baseResetTokenService.GenerateToken()

	assert.NotEmpty(t, uuidOne)
	assert.NotEmpty(t, uuidTwo)

	assert.NotEqual(t, uuidOne, uuidTwo)

	_, err := uuid.Parse(uuidOne)

	assert.Nil(t, err)

	_, err = uuid.Parse(uuidTwo)

	assert.Nil(t, err)
}

func TestBaseResetTokenService_ResetPasswordByTokenAndNewPassword(t *testing.T) {
	mockUserService := MockUserService{}
	mockUserRepository := MockUserRepository{}
	mockResetTokenRepository := MockResetTokenRepository{}

	mockResetToken := &repositories.GormResetToken{}
	mockResetTokenRepository.On("GetResetTokenByToken", "1").Return(mockResetToken)
	mockResetTokenRepository.On("GetResetTokenByToken", "2").Return(nil)
	mockUserService.On("HashPassword", "1").Return("1", nil)
	mockUserService.On("HashPassword", "2").Return("", errors.New("invalid-password"))
	mockUserRepository.On("UpdatePassword", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockResetTokenRepository.On("DeleteResetToken", mock.Anything).Return(nil)

	baseResetTokenService := NewBaseResetTokenService(&mockUserService, &mockResetTokenRepository, &mockUserRepository)

	tests := []struct {
		name         string
		token        string
		password     string
		expectError  bool
		errorMessage string
	}{
		{
			name:         "Valid token and password",
			token:        "1",
			password:     "1",
			expectError:  false,
			errorMessage: "",
		},
		{
			name:         "Token not found",
			token:        "2",
			password:     "1",
			expectError:  true,
			errorMessage: "token not found",
		},
		{
			name:         "Invalid password for hashing",
			token:        "1",
			password:     "2",
			expectError:  true,
			errorMessage: "invalid-password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := baseResetTokenService.ResetPasswordByTokenAndNewPassword(tt.token, tt.password)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMessage, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestBaseResetTokenService_ResetPasswordByUserUUIDAndOldPasswordAndNewPassword(t *testing.T) {
	mockUserService := MockUserService{}
	mockUserRepository := MockUserRepository{}
	mockResetTokenRepository := MockResetTokenRepository{}

	mockUser := &repositories.GormUser{}
	mockUserRepository.On("GetUserByUUIDAndPassword", "1", mock.Anything).Return(mockUser)
	mockUserRepository.On("GetUserByUUIDAndPassword", "2", mock.Anything).Return(nil)
	mockUserService.On("HashPassword", "1").Return("1", nil)
	mockUserService.On("HashPassword", "2").Return("", errors.New("invalid-password"))
	mockUserRepository.On("UpdatePassword", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockResetTokenRepository.On("DeleteResetToken", mock.Anything).Return(nil)

	baseResetTokenService := NewBaseResetTokenService(&mockUserService, &mockResetTokenRepository, &mockUserRepository)

	tests := []struct {
		name         string
		userUUID     string
		oldPassword  string
		newPassword  string
		expectError  bool
		errorMessage string
	}{
		{
			name:         "Valid userUUID, oldPassword and newPassword",
			userUUID:     "1",
			oldPassword:  "1",
			newPassword:  "1",
			expectError:  false,
			errorMessage: "",
		},
		{
			name:         "UserUUID not found",
			userUUID:     "2",
			oldPassword:  "1",
			newPassword:  "1",
			expectError:  true,
			errorMessage: "user not found",
		},
		{
			name:         "Invalid password for hashing",
			userUUID:     "1",
			oldPassword:  "1",
			newPassword:  "2",
			expectError:  true,
			errorMessage: "invalid-password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := baseResetTokenService.ResetPasswordByUserUUIDAndOldPasswordAndNewPassword(tt.userUUID, tt.oldPassword, tt.newPassword)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMessage, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
