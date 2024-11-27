package services

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	services_mocks "github.com/fromsi/jwt-oauth-sso/mocks/services"
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
	mockResetTokenBuilder := repositories_mocks.NewMockResetTokenBuilder(mockController)

	baseResetTokenService := NewBaseResetTokenService(
		config,
		mockUserService,
		mockResetTokenRepository,
		mockUserRepository,
		mockNotificationService,
		mockResetTokenBuilder,
	)

	assert.NotEmpty(t, baseResetTokenService)
}

func TestBaseResetTokenService_GenerateToken(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockResetTokenRepository := repositories_mocks.NewMockResetTokenRepository(mockController)
	mockNotificationService := services_mocks.NewMockNotificationService(mockController)
	mockResetTokenBuilder := repositories_mocks.NewMockResetTokenBuilder(mockController)
	mockResetToken := repositories_mocks.NewMockResetToken(mockController)

	mockResetTokenBuilder.EXPECT().New().Return(mockResetTokenBuilder).AnyTimes()
	mockResetTokenBuilder.EXPECT().SetToken(gomock.Any()).Return(mockResetTokenBuilder).AnyTimes()
	mockResetTokenBuilder.EXPECT().SetUserUUID(gomock.Any()).Return(mockResetTokenBuilder).AnyTimes()
	mockResetTokenBuilder.EXPECT().SetExpiresAt(gomock.Any()).Return(mockResetTokenBuilder).AnyTimes()
	mockResetTokenBuilder.EXPECT().SetCreatedAt(gomock.Any()).Return(mockResetTokenBuilder).AnyTimes()
	mockResetTokenBuilder.EXPECT().Build().Return(mockResetToken, nil).AnyTimes()

	baseResetTokenService := NewBaseResetTokenService(
		config,
		mockUserService,
		mockResetTokenRepository,
		mockUserRepository,
		mockNotificationService,
		mockResetTokenBuilder,
	)

	uuidOne := baseResetTokenService.GenerateToken()
	uuidTwo := baseResetTokenService.GenerateToken()

	assert.NotEmpty(t, uuidOne)
	assert.NotEmpty(t, uuidTwo)

	assert.NotEqual(t, uuidOne, uuidTwo)

	_, err := uuid.Parse(uuidOne)

	assert.NoError(t, err)

	_, err = uuid.Parse(uuidTwo)

	assert.NoError(t, err)
}

func TestBaseResetTokenService_SendNewResetTokenByUser(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockResetTokenRepository := repositories_mocks.NewMockResetTokenRepository(mockController)
	mockNotificationService := services_mocks.NewMockNotificationService(mockController)
	mockResetTokenBuilder := repositories_mocks.NewMockResetTokenBuilder(mockController)
	mockResetToken := repositories_mocks.NewMockResetToken(mockController)
	mockUser := repositories_mocks.NewMockUser(mockController)

	mockUser.EXPECT().GetUUID().Return("1").AnyTimes()

	mockResetToken.EXPECT().GetToken().Return("1").AnyTimes()

	mockResetTokenBuilder.EXPECT().New().Return(mockResetTokenBuilder).AnyTimes()
	mockResetTokenBuilder.EXPECT().SetToken(gomock.Any()).Return(mockResetTokenBuilder).AnyTimes()
	mockResetTokenBuilder.EXPECT().SetUserUUID(gomock.Any()).Return(mockResetTokenBuilder).AnyTimes()
	mockResetTokenBuilder.EXPECT().SetExpiresAt(gomock.Any()).Return(mockResetTokenBuilder).AnyTimes()
	mockResetTokenBuilder.EXPECT().SetCreatedAt(gomock.Any()).Return(mockResetTokenBuilder).AnyTimes()
	mockResetTokenBuilder.EXPECT().Build().Return(mockResetToken, nil)

	mockResetTokenRepository.EXPECT().CreateResetToken(gomock.Any()).Return(nil)

	mockNotificationService.EXPECT().SendTextByUser(gomock.Any(), gomock.Any()).Return(nil)

	baseResetTokenService := NewBaseResetTokenService(
		config,
		mockUserService,
		mockResetTokenRepository,
		mockUserRepository,
		mockNotificationService,
		mockResetTokenBuilder,
	)

	err := baseResetTokenService.SendNewResetTokenByUser(mockUser)

	assert.NoError(t, err)

	mockResetTokenBuilder.EXPECT().Build().Return(mockResetToken, nil)
	mockResetTokenRepository.EXPECT().CreateResetToken(gomock.Any()).Return(nil)
	mockNotificationService.EXPECT().SendTextByUser(gomock.Any(), gomock.Any()).Return(errors.New("error"))

	err = baseResetTokenService.SendNewResetTokenByUser(mockUser)

	assert.Error(t, err)

	mockResetTokenBuilder.EXPECT().Build().Return(mockResetToken, nil)
	mockResetTokenRepository.EXPECT().CreateResetToken(gomock.Any()).Return(errors.New("error"))

	err = baseResetTokenService.SendNewResetTokenByUser(mockUser)

	assert.Error(t, err)

	mockResetTokenBuilder.EXPECT().Build().Return(nil, errors.New("error"))

	err = baseResetTokenService.SendNewResetTokenByUser(mockUser)

	assert.Error(t, err)
}
