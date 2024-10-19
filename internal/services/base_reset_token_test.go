package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewBaseResetTokenService(t *testing.T) {
	baseResetTokenService := NewBaseResetTokenService()

	assert.NotNil(t, baseResetTokenService)
}

func TestBaseResetTokenService_GenerateToken(t *testing.T) {
	baseResetTokenService := NewBaseResetTokenService()

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
