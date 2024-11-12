package services

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
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
	baseUserService := NewBaseUserService(mockUserRepository)

	assert.NotEmpty(t, baseUserService)
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

	userOne := repositories.NewGormUser()

	userOne.SetUUID("1")
	userOne.SetEmail("1")
	userOne.SetPassword("1")

	userTwo := repositories.NewGormUser()

	userTwo.SetUUID("2")
	userTwo.SetEmail("2")
	userTwo.SetPassword("2")

	mockUserRepository.
		EXPECT().
		CreateUser(gomock.Any()).
		DoAndReturn(func(user repositories.User) error {
			isEqualUUID := user.GetUUID() == userTwo.GetUUID()
			isEqualEmail := user.GetEmail() == userTwo.GetEmail()
			isEqualPassword := user.GetPassword() == userTwo.GetPassword()

			if isEqualUUID && isEqualEmail && isEqualPassword {
				return errors.New("error")
			}

			return nil
		}).
		AnyTimes()

	err := baseUserService.CreateUserByUUIDAndEmailAndHashedPassword(
		userOne.GetUUID(),
		userOne.GetEmail(),
		userOne.GetPassword(),
	)

	assert.NoError(t, err)

	err = baseUserService.CreateUserByUUIDAndEmailAndHashedPassword(
		userTwo.GetUUID(),
		userTwo.GetEmail(),
		userTwo.GetPassword(),
	)

	assert.Error(t, err)
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
