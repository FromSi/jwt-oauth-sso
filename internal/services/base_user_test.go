package services

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func Test_NewBaseUserService(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	baseUserService := NewBaseUserService(mockUserRepository)

	assert.NotNil(t, baseUserService)
}

func TestBaseUserService_GenerateUUID(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	baseUserService := NewBaseUserService(mockUserRepository)

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
	baseUserService := NewBaseUserService(mockUserRepository)

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
}

func TestBaseUserService_CheckHashedPasswordAndNativePassword(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	baseUserService := NewBaseUserService(mockUserRepository)

	hashedPassword, err := baseUserService.HashPassword("1")

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = baseUserService.CheckHashedPasswordAndNativePassword(hashedPassword, "1")

	assert.NoError(t, err)

	err = baseUserService.CheckHashedPasswordAndNativePassword(hashedPassword, "2")

	assert.Error(t, err)
}

func TestBaseUserService_CreateUserByUUIDAndEmailAndHashedPassword(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	baseUserService := NewBaseUserService(mockUserRepository)

	user := repositories.NewGormUser()

	user.SetUUID("1")
	user.SetEmail("1")
	user.SetPassword("1")

	mockUserRepository.
		EXPECT().
		CreateUser(gomock.Any()).
		DoAndReturn(func(user repositories.User) error {
			assert.Equal(t, "1", user.GetUUID())
			assert.Equal(t, "1", user.GetEmail())
			assert.Equal(t, "1", user.GetPassword())

			return nil
		}).
		AnyTimes()

	err := baseUserService.CreateUserByUUIDAndEmailAndHashedPassword(
		user.GetUUID(),
		user.GetEmail(),
		user.GetPassword(),
	)

	assert.NoError(t, err)
}

func TestBaseUserService_UpdatePasswordByUUIDAndHashedPassword(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	baseUserService := NewBaseUserService(mockUserRepository)

	mockUserRepository.
		EXPECT().
		UpdatePasswordByUUIDAndPasswordAndUpdatedAt("1", gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	mockUserRepository.
		EXPECT().
		UpdatePasswordByUUIDAndPasswordAndUpdatedAt(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(errors.New("error")).
		AnyTimes()

	err := baseUserService.UpdatePasswordByUUIDAndHashedPassword("1", "2")

	assert.NoError(t, err)

	err = baseUserService.UpdatePasswordByUUIDAndHashedPassword("0", "0")

	assert.Error(t, err)
}
