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

func (receiver *BaseDeviceService) GetDeviceByUserUUIDAndIpAndAgent(
	config configs.TokenConfig,
	userUUID string,
	ip string,
	agent string,
) repositories.Device {
	device := receiver.deviceRepository.GetDeviceByUserUUIDAndIpAndAgent(userUUID, ip, agent)

	if device != nil {
		return device
	}

	device = repositories.NewGormDevice()

	device.SetUUID(receiver.GenerateUUID())
	device.SetUserUUID(userUUID)
	device.SetIp(ip)
	device.SetAgent(agent)

	timeNow := time.Now()

	device.SetCreatedAt(int(timeNow.Unix()))
	device.SetUpdatedAt(int(timeNow.Unix()))
	device.SetExpiredAt(int(timeNow.AddDate(0, 0, config.GetExpirationRefreshInDays()).Unix()))

	return device
}
