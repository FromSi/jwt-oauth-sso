package services

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"testing"
)

func Test_NewBaseUserService(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockUserBuilder := repositories_mocks.NewMockUserBuilder(mockController)

	baseUserService := NewBaseUserService(mockUserRepository, mockUserBuilder)

	assert.NotEmpty(t, baseUserService)
}

func TestBaseUserService_GenerateUUID(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockUserBuilder := repositories_mocks.NewMockUserBuilder(mockController)

	baseUserService := NewBaseUserService(mockUserRepository, mockUserBuilder)

	uuidOne := baseUserService.GenerateUUID()
	uuidTwo := baseUserService.GenerateUUID()

	assert.NotEmpty(t, uuidOne)
	assert.NotEmpty(t, uuidTwo)

	assert.NotEqual(t, uuidOne, uuidTwo)

	_, err := uuid.Parse(uuidOne)

	assert.NoError(t, err)

	_, err = uuid.Parse(uuidTwo)

	assert.NoError(t, err)
}

func TestBaseUserService_HashPassword(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockUserBuilder := repositories_mocks.NewMockUserBuilder(mockController)

	baseUserService := NewBaseUserService(mockUserRepository, mockUserBuilder)

	hashedPassword, err := baseUserService.HashPassword("1")

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte("1"))

	assert.NoError(t, err)

	hashedPassword, err = baseUserService.HashPassword("2")

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte("2"))

	assert.NoError(t, err)

	var passwordBuilder strings.Builder

	for i := 0; i < 100; i++ {
		passwordBuilder.WriteString("1")
	}

	hashedPassword, err = baseUserService.HashPassword(passwordBuilder.String())

	assert.Error(t, err)
	assert.Empty(t, hashedPassword)
}

func TestBaseUserService_CheckHashedPasswordAndNativePassword(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockUserBuilder := repositories_mocks.NewMockUserBuilder(mockController)

	baseUserService := NewBaseUserService(mockUserRepository, mockUserBuilder)

	hashedPassword, err := baseUserService.HashPassword("1")

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = baseUserService.CheckHashedPasswordAndNativePassword(hashedPassword, "1")

	assert.NoError(t, err)

	hashedPassword, err = baseUserService.HashPassword("2")

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = baseUserService.CheckHashedPasswordAndNativePassword(hashedPassword, "1")

	assert.Error(t, err)
}

func TestBaseUserService_CreateUserByUUIDAndEmailAndHashedPassword(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockUserBuilder := repositories_mocks.NewMockUserBuilder(mockController)
	mockUser := repositories_mocks.NewMockUser(mockController)

	mockUserBuilder.EXPECT().New().Return(mockUserBuilder)
	mockUserBuilder.EXPECT().SetUUID(gomock.Any()).Return(mockUserBuilder)
	mockUserBuilder.EXPECT().SetEmail(gomock.Any()).Return(mockUserBuilder)
	mockUserBuilder.EXPECT().SetPassword(gomock.Any()).Return(mockUserBuilder)
	mockUserBuilder.EXPECT().SetCreatedAt(gomock.Any()).Return(mockUserBuilder)
	mockUserBuilder.EXPECT().SetUpdatedAt(gomock.Any()).Return(mockUserBuilder)
	mockUserBuilder.EXPECT().Build().Return(mockUser, nil)

	mockUserRepository.EXPECT().CreateUser(gomock.Any()).Return(nil)

	baseUserService := NewBaseUserService(mockUserRepository, mockUserBuilder)

	err := baseUserService.CreateUserByUUIDAndEmailAndHashedPassword("1", "1", "1")

	assert.NoError(t, err)

	mockUserBuilder.EXPECT().New().Return(mockUserBuilder)
	mockUserBuilder.EXPECT().SetUUID(gomock.Any()).Return(mockUserBuilder)
	mockUserBuilder.EXPECT().SetEmail(gomock.Any()).Return(mockUserBuilder)
	mockUserBuilder.EXPECT().SetPassword(gomock.Any()).Return(mockUserBuilder)
	mockUserBuilder.EXPECT().SetCreatedAt(gomock.Any()).Return(mockUserBuilder)
	mockUserBuilder.EXPECT().SetUpdatedAt(gomock.Any()).Return(mockUserBuilder)
	mockUserRepository.EXPECT().CreateUser(gomock.Any()).Return(errors.New("error"))
	mockUserBuilder.EXPECT().Build().Return(mockUser, nil)

	err = baseUserService.CreateUserByUUIDAndEmailAndHashedPassword("1", "1", "1")

	assert.Error(t, err)

	mockUserBuilder.EXPECT().New().Return(mockUserBuilder)
	mockUserBuilder.EXPECT().SetUUID(gomock.Any()).Return(mockUserBuilder)
	mockUserBuilder.EXPECT().SetEmail(gomock.Any()).Return(mockUserBuilder)
	mockUserBuilder.EXPECT().SetPassword(gomock.Any()).Return(mockUserBuilder)
	mockUserBuilder.EXPECT().SetCreatedAt(gomock.Any()).Return(mockUserBuilder)
	mockUserBuilder.EXPECT().SetUpdatedAt(gomock.Any()).Return(mockUserBuilder)
	mockUserBuilder.EXPECT().Build().Return(nil, errors.New("error"))

	err = baseUserService.CreateUserByUUIDAndEmailAndHashedPassword("1", "1", "1")

	assert.Error(t, err)
}

func TestBaseUserService_UpdatePasswordByUUIDAndHashedPassword(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockUserBuilder := repositories_mocks.NewMockUserBuilder(mockController)

	baseUserService := NewBaseUserService(mockUserRepository, mockUserBuilder)

	mockUserRepository.
		EXPECT().
		UpdatePasswordByUUIDAndPasswordAndUpdatedAt(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)

	err := baseUserService.UpdatePasswordByUUIDAndHashedPassword("1", "2")

	assert.NoError(t, err)

	mockUserRepository.
		EXPECT().
		UpdatePasswordByUUIDAndPasswordAndUpdatedAt(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(errors.New("error"))

	err = baseUserService.UpdatePasswordByUUIDAndHashedPassword("0", "0")

	assert.Error(t, err)
}
