package repositories

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
)

type QueryDeviceRepository interface {
	GetDevicesByUserUUID(string) []Device
	GetDeviceByUserUUIDAndIpAndAgent(string, string, string) Device
}

type MutableDeviceRepository interface {
	CreateDevice(Device) error
	UpdateDevice(Device) error
	DeleteDeviceByUUID(string) error
	DeleteAllDevicesByUserUUID(string) error
}

type DeviceRepository interface {
	QueryDeviceRepository
	MutableDeviceRepository
}

type QueryDevice interface {
	GenerateAccessToken(configs.TokenConfig) (*tokens.AccessToken, error)
	GetUUID() string
	GetUserUUID() string
	GetAgent() string
	GetIp() string
	GetRefreshToken() string
	GetExpiredAt() int
	GetCreatedAt() int
	GetUpdatedAt() int
}

type MutableDevice interface {
	SetUUID(string)
	SetUserUUID(string)
	SetAgent(string)
	SetIp(string)
	SetRefreshToken(string)
	SetExpiredAt(int)
	SetCreatedAt(int)
	SetUpdatedAt(int)
}

type Device interface {
	QueryDevice
	MutableDevice
}
