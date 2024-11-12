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
}

func NewBaseDeviceService(
	config configs.TokenConfig,
	deviceRepository repositories.DeviceRepository,
) *BaseDeviceService {
	return &BaseDeviceService{
		config:           config,
		deviceRepository: deviceRepository,
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
) repositories.Device {
	timeNow := time.Now()
	device := repositories.NewGormDevice()

	device.SetUUID(receiver.GenerateUUID())
	device.SetUserUUID(userUUID)
	device.SetIp(ip)
	device.SetUserAgent(userAgent)

	device.SetCreatedAt(int(timeNow.Unix()))
	device.SetUpdatedAt(int(timeNow.Unix()))

	return device
}

func (receiver *BaseDeviceService) GetNewRefreshDetailsByDevice(
	device repositories.Device,
) repositories.Device {
	timeNow := time.Now()

	device.SetRefreshToken(receiver.GenerateRefreshToken())
	device.SetUpdatedAt(int(timeNow.Unix()))

	expiresAt := timeNow.
		AddDate(0, 0, receiver.config.GetExpirationRefreshInDays()).
		Unix()

	device.SetExpiresAt(int(expiresAt))

	return device
}
