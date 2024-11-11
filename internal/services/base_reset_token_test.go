package services

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/internal/mocks/repositories"
	services_mocks "github.com/fromsi/jwt-oauth-sso/internal/mocks/services"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewBaseResetTokenService(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockResetTokenRepository := repositories_mocks.NewMockResetTokenRepository(mockController)
	mockNotificationService := services_mocks.NewMockNotificationService(mockController)

	baseResetTokenService := NewBaseResetTokenService(
		config,
		mockUserService,
		mockResetTokenRepository,
		mockUserRepository,
		mockNotificationService,
	)

	assert.NotNil(t, baseResetTokenService)
}

func TestBaseResetTokenService_GenerateToken(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockResetTokenRepository := repositories_mocks.NewMockResetTokenRepository(mockController)
	mockNotificationService := services_mocks.NewMockNotificationService(mockController)

	baseResetTokenService := NewBaseResetTokenService(
		config,
		mockUserService,
		mockResetTokenRepository,
		mockUserRepository,
		mockNotificationService,
	)

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

func TestBaseResetTokenService_SendNewResetTokenByUserEmail(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockResetTokenRepository := repositories_mocks.NewMockResetTokenRepository(mockController)
	mockNotificationService := services_mocks.NewMockNotificationService(mockController)

	user := repositories.NewGormUser()
	mockUserRepository.
		EXPECT().
		GetUserByEmail(gomock.Eq("1")).
		Return(user).
		AnyTimes()

	mockUserRepository.
		EXPECT().
		GetUserByEmail(gomock.Any()).
		Return(nil).
		AnyTimes()

	mockResetTokenRepository.
		EXPECT().
		CreateResetToken(gomock.Any()).
		Return(nil).
		AnyTimes()

	mockNotificationService.
		EXPECT().
		SendTextByUser(gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	baseResetTokenService := NewBaseResetTokenService(
		config,
		mockUserService,
		mockResetTokenRepository,
		mockUserRepository,
		mockNotificationService,
	)

	err := baseResetTokenService.SendNewResetTokenByUserEmail("1")

	assert.NoError(t, err)

	err = baseResetTokenService.SendNewResetTokenByUserEmail("2")

	assert.Error(t, err)
}

func TestBaseResetTokenService_ResetPasswordByUserUUIDAndNewPassword(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockResetTokenRepository := repositories_mocks.NewMockResetTokenRepository(mockController)
	mockNotificationService := services_mocks.NewMockNotificationService(mockController)

	mockResetToken := repositories.NewGormResetToken()
	mockResetTokenRepository.
		EXPECT().
		GetActiveResetTokenByToken(gomock.Eq("1")).
		Return(mockResetToken).
		AnyTimes()

	mockResetTokenRepository.
		EXPECT().
		GetActiveResetTokenByToken(gomock.Eq("2")).
		Return(nil).
		AnyTimes()

	mockUserService.
		EXPECT().
		HashPassword(gomock.Eq("1")).
		Return("1", nil).
		AnyTimes()

	mockUserService.
		EXPECT().
		HashPassword(gomock.Eq("2")).
		Return("", errors.New("invalid-password")).
		AnyTimes()

	mockUserRepository.
		EXPECT().
		UpdatePasswordByUUIDAndPasswordAndUpdatedAt(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	baseResetTokenService := NewBaseResetTokenService(
		config,
		mockUserService,
		mockResetTokenRepository,
		mockUserRepository,
		mockNotificationService,
	)

	tests := []struct {
		name         string
		userUUID     string
		password     string
		expectError  bool
		errorMessage string
	}{
		{
			name:         "Valid token and password",
			userUUID:     "1",
			password:     "1",
			expectError:  false,
			errorMessage: "",
		},
		{
			name:         "Invalid password for hashing",
			userUUID:     "1",
			password:     "2",
			expectError:  true,
			errorMessage: "invalid-password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := baseResetTokenService.
				ResetPasswordByUserUUIDAndNewPassword(tt.userUUID, tt.password)

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
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockResetTokenRepository := repositories_mocks.NewMockResetTokenRepository(mockController)
	mockNotificationService := services_mocks.NewMockNotificationService(mockController)

	mockUserRepository.
		EXPECT().
		HasUserByUUIDAndPassword(gomock.Eq("1"), gomock.Any()).
		Return(true).
		AnyTimes()

	mockUserRepository.
		EXPECT().
		HasUserByUUIDAndPassword(gomock.Eq("2"), gomock.Any()).
		Return(false).
		AnyTimes()

	mockUserService.
		EXPECT().
		HashPassword(gomock.Eq("1")).
		Return("1", nil).
		AnyTimes()

	mockUserService.
		EXPECT().
		HashPassword(gomock.Eq("2")).
		Return("", errors.New("invalid-password")).
		AnyTimes()

	mockUserRepository.
		EXPECT().
		UpdatePasswordByUUIDAndPasswordAndUpdatedAt(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	mockResetTokenRepository.
		EXPECT().
		DeleteResetToken(gomock.Any()).
		Return(nil).
		AnyTimes()

	baseResetTokenService := NewBaseResetTokenService(
		config,
		mockUserService,
		mockResetTokenRepository,
		mockUserRepository,
		mockNotificationService,
	)

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
			err := baseResetTokenService.
				ResetPasswordByUserUUIDAndOldPasswordAndNewPassword(
					tt.userUUID,
					tt.oldPassword,
					tt.newPassword,
				)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMessage, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
