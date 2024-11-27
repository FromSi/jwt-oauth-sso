package services

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/google/uuid"
	"time"
)

type BaseDeviceService struct {
	config           configs.TokenConfig
	deviceRepository repositories.DeviceRepository
	deviceBuilder    repositories.DeviceBuilder
}

func NewBaseDeviceService(
	config configs.TokenConfig,
	deviceRepository repositories.DeviceRepository,
	deviceBuilder repositories.DeviceBuilder,
) *BaseDeviceService {
	return &BaseDeviceService{
		config:           config,
		deviceRepository: deviceRepository,
		deviceBuilder:    deviceBuilder,
	}
}

func (receiver *BaseDeviceService) GenerateUUID() string {
	return uuid.New().String()
}

func (receiver *BaseDeviceService) GenerateRefreshToken() string {
	return uuid.New().String()
}

func (receiver *BaseDeviceService) GetOldDeviceByUserUUIDAndIpAndUserAgent(
	userUUID string,
	ip string,
	userAgent string,
) repositories.Device {
	return receiver.
		deviceRepository.
		GetDeviceByUserUUIDAndIpAndUserAgent(userUUID, ip, userAgent)
}

func (receiver *BaseDeviceService) GetNewDeviceByUserUUIDAndIpAndUserAgent(
	userUUID string,
	ip string,
	userAgent string,
) (repositories.Device, error) {
	timeNow := time.Now()

	return receiver.deviceBuilder.
		New().
		SetUUID(receiver.GenerateUUID()).
		SetUserUUID(userUUID).
		SetIp(ip).
		SetUserAgent(userAgent).
		SetCreatedAt(int(timeNow.Unix())).
		SetUpdatedAt(int(timeNow.Unix())).
		Build()
}

func (receiver *BaseDeviceService) GetNewRefreshDetailsByDevice(
	device repositories.Device,
) (repositories.Device, error) {
	timeNow := time.Now()
	expiresAt := timeNow.
		AddDate(0, 0, receiver.config.GetExpirationRefreshInDays()).
		Unix()

	deviceResult, err := receiver.
		deviceBuilder.
		NewFromDevice(device).
		SetRefreshToken(receiver.GenerateRefreshToken()).
		SetUpdatedAt(int(timeNow.Unix())).
		SetExpiresAt(int(expiresAt)).
		Build()

	if err != nil {
		return nil, err
	}

	return deviceResult, nil
}
