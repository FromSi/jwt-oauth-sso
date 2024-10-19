package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewBaseDeviceService(t *testing.T) {
	baseDeviceService := NewBaseDeviceService()

	assert.NotNil(t, baseDeviceService)
}

func TestBaseDeviceService_GenerateUUID(t *testing.T) {
	baseDeviceService := NewBaseDeviceService()

	uuidOne := baseDeviceService.GenerateUUID()
	uuidTwo := baseDeviceService.GenerateUUID()

	assert.NotEmpty(t, uuidOne)
	assert.NotEmpty(t, uuidTwo)

	assert.NotEqual(t, uuidOne, uuidTwo)

	_, err := uuid.Parse(uuidOne)

	assert.Nil(t, err)

	_, err = uuid.Parse(uuidTwo)

	assert.Nil(t, err)
}
