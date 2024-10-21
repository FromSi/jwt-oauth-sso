package services

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

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

	mockUserRepository.On("HasUserByUUIDAndPassword", "1", mock.Anything).Return(true)
	mockUserRepository.On("HasUserByUUIDAndPassword", "2", mock.Anything).Return(false)
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
