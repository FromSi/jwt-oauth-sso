package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func Test_NewBaseUserService(t *testing.T) {
	baseUserService := NewBaseUserService()

	assert.NotNil(t, baseUserService)
}

func TestBaseUserService_GenerateUUID(t *testing.T) {
	baseUserService := NewBaseUserService()

	uuidOne := baseUserService.GenerateUUID()
	uuidTwo := baseUserService.GenerateUUID()

	assert.NotEmpty(t, uuidOne)
	assert.NotEmpty(t, uuidTwo)

	assert.NotEqual(t, uuidOne, uuidTwo)

	_, err := uuid.Parse(uuidOne)

	assert.Nil(t, err)

	_, err = uuid.Parse(uuidTwo)

	assert.Nil(t, err)
}

func TestBaseUserService_HashPassword(t *testing.T) {
	baseUserService := NewBaseUserService()

	hashedPassword, err := baseUserService.HashPassword("1")

	assert.Nil(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte("1"))

	assert.Nil(t, err)
}

func TestBaseUserService_CheckPasswordByHashAndPassword(t *testing.T) {
	baseUserService := NewBaseUserService()

	hashedPassword, err := baseUserService.HashPassword("1")

	assert.Nil(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = baseUserService.CheckPasswordByHashAndPassword(hashedPassword, "1")

	assert.Nil(t, err)

	err = baseUserService.CheckPasswordByHashAndPassword(hashedPassword, "2")

	assert.NotNil(t, err)
}
