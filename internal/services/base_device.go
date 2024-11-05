package services

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/google/uuid"
	"time"
)

type BaseDeviceService struct {
	deviceRepository repositories.DeviceRepository
}

func NewBaseDeviceService(deviceRepository repositories.DeviceRepository) *BaseDeviceService {
	return &BaseDeviceService{deviceRepository: deviceRepository}
}

func (receiver *BaseDeviceService) GenerateUUID() string {
	return uuid.New().String()
}

func (receiver *BaseDeviceService) GenerateRefreshToken() string {
	return uuid.New().String()
}

func (receiver *BaseDeviceService) GetDeviceByUserUUIDAndIpAndUserAgent(
	config configs.TokenConfig,
	userUUID string,
	ip string,
	userAgent string,
) (repositories.Device, error) {
	timeNow := time.Now()
	device := receiver.deviceRepository.GetDeviceByUserUUIDAndIpAndUserAgent(userUUID, ip, userAgent)

	if device == nil {
		return nil, nil
	}

	device.SetRefreshToken(receiver.GenerateRefreshToken())
	device.SetUpdatedAt(int(timeNow.Unix()))
	device.SetExpiredAt(int(timeNow.AddDate(0, 0, config.GetExpirationRefreshInDays()).Unix()))

	err := receiver.deviceRepository.UpdateDevice(device)

	if err != nil {
		return nil, err
	}

	return device, nil
}

func (receiver *BaseDeviceService) GetNewDeviceByUserUUIDAndIpAndUserAgent(
	config configs.TokenConfig,
	userUUID string,
	ip string,
	userAgent string,
) (repositories.Device, error) {
	timeNow := time.Now()
	device := repositories.NewGormDevice()

	device.SetUUID(receiver.GenerateUUID())
	device.SetUserUUID(userUUID)
	device.SetIp(ip)
	device.SetUserAgent(userAgent)
	device.SetRefreshToken(receiver.GenerateRefreshToken())

	device.SetCreatedAt(int(timeNow.Unix()))
	device.SetUpdatedAt(int(timeNow.Unix()))
	device.SetExpiredAt(int(timeNow.AddDate(0, 0, config.GetExpirationRefreshInDays()).Unix()))

	err := receiver.deviceRepository.CreateDevice(device)

	if err != nil {
		return nil, err
	}

	return device, nil
}

func (receiver *BaseDeviceService) ResetDevice(
	config configs.TokenConfig,
	device repositories.Device,
) (repositories.Device, error) {
	timeNow := time.Now()

	deviceForUpdate := repositories.NewGormDeviceByDevice(device)

	deviceForUpdate.SetRefreshToken(receiver.GenerateRefreshToken())
	deviceForUpdate.SetUpdatedAt(int(timeNow.Unix()))
	deviceForUpdate.SetExpiredAt(int(timeNow.AddDate(0, 0, config.GetExpirationRefreshInDays()).Unix()))

	err := receiver.deviceRepository.UpdateDevice(deviceForUpdate)

	if err != nil {
		return nil, err
	}

	return deviceForUpdate, nil
}
