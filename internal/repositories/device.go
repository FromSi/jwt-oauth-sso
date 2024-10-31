package repositories

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
)

//go:generate mockgen -destination=../mocks/repositories/mock_query_device_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories QueryDeviceRepository
type QueryDeviceRepository interface {
	GetDevicesByUserUUID(string) []Device
	GetDeviceByUserUUIDAndIpAndAgent(string, string, string) Device
}

//go:generate mockgen -destination=../mocks/repositories/mock_mutable_device_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories MutableDeviceRepository
type MutableDeviceRepository interface {
	CreateDevice(Device) error
	UpdateDevice(Device) error
	DeleteDeviceByUUID(string) error
	DeleteAllDevicesByUserUUID(string) error
}

//go:generate mockgen -destination=../mocks/repositories/mock_device_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories DeviceRepository
type DeviceRepository interface {
	QueryDeviceRepository
	MutableDeviceRepository
}

//go:generate mockgen -destination=../mocks/repositories/mock_query_device.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories QueryDevice
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

//go:generate mockgen -destination=../mocks/repositories/mock_mutable_device.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories MutableDevice
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

//go:generate mockgen -destination=../mocks/repositories/mock_device.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories Device
type Device interface {
	QueryDevice
	MutableDevice
}
