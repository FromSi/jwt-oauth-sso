package services

import "github.com/google/uuid"

type BaseDeviceService struct{}

func NewBaseDeviceService() *BaseDeviceService {
	return &BaseDeviceService{}
}

func (receiver BaseDeviceService) GenerateUUID() string {
	return uuid.New().String()
}
