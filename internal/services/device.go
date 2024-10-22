package services

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
)

type QueryDeviceService interface {
	GenerateUUID() string
	GenerateRefreshToken() string
	GetDeviceByUserUUIDAndIpAndAgent(configs.TokenConfig, string, string, string) repositories.Device
}

type MutableDeviceService interface {
}

type DeviceService interface {
	QueryDeviceService
	MutableDeviceService
}
