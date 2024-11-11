package repositories

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
)

//go:generate mockgen -destination=../../mocks/repositories/mock_query_device_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories QueryDeviceRepository
type QueryDeviceRepository interface {
	GetDeviceByRefreshToken(string) Device
	GetDevicesByUserUUID(string) []Device
	GetDeviceByUserUUIDAndIpAndUserAgent(string, string, string) Device
}

//go:generate mockgen -destination=../../mocks/repositories/mock_mutable_device_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories MutableDeviceRepository
type MutableDeviceRepository interface {
	CreateDevice(Device) error
	UpdateDevice(Device) error
	DeleteDeviceByUUID(string) error
	DeleteDeviceByUUIDAndUserUUID(string, string) error
	DeleteAllDevicesByUserUUID(string) error
}

//go:generate mockgen -destination=../../mocks/repositories/mock_device_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories DeviceRepository
type DeviceRepository interface {
	QueryDeviceRepository
	MutableDeviceRepository
}

//go:generate mockgen -destination=../../mocks/repositories/mock_query_device.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories QueryDevice
type QueryDevice interface {
	GenerateAccessToken(configs.TokenConfig) (*tokens.AccessToken, error)
	GetUUID() string
	GetUserUUID() string
	GetUserAgent() string
	GetIp() string
	GetRefreshToken() string
	GetExpiresAt() int
	GetCreatedAt() int
	GetUpdatedAt() int
}

//go:generate mockgen -destination=../../mocks/repositories/mock_mutable_device.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories MutableDevice
type MutableDevice interface {
	SetUUID(string)
	SetUserUUID(string)
	SetUserAgent(string)
	SetIp(string)
	SetRefreshToken(string)
	SetExpiresAt(int)
	SetCreatedAt(int)
	SetUpdatedAt(int)
}

//go:generate mockgen -destination=../../mocks/repositories/mock_device.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories Device
type Device interface {
	QueryDevice
	MutableDevice
}
