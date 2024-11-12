package services

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	repositories_mocks2 "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	services_mocks2 "github.com/fromsi/jwt-oauth-sso/mocks/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewBaseResetTokenService(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockUserService := services_mocks2.NewMockUserService(mockController)
	mockUserRepository := repositories_mocks2.NewMockUserRepository(mockController)
	mockResetTokenRepository := repositories_mocks2.NewMockResetTokenRepository(mockController)
	mockNotificationService := services_mocks2.NewMockNotificationService(mockController)

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
	mockUserService := services_mocks2.NewMockUserService(mockController)
	mockUserRepository := repositories_mocks2.NewMockUserRepository(mockController)
	mockResetTokenRepository := repositories_mocks2.NewMockResetTokenRepository(mockController)
	mockNotificationService := services_mocks2.NewMockNotificationService(mockController)

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

func TestBaseResetTokenService_SendNewResetTokenByUser(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockUserService := services_mocks2.NewMockUserService(mockController)
	mockUserRepository := repositories_mocks2.NewMockUserRepository(mockController)
	mockResetTokenRepository := repositories_mocks2.NewMockResetTokenRepository(mockController)
	mockNotificationService := services_mocks2.NewMockNotificationService(mockController)

	user := repositories.NewGormUser()

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

	err := baseResetTokenService.SendNewResetTokenByUser(user)

	assert.NoError(t, err)
}
