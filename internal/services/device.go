package services

import (
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
)

//go:generate mockgen -destination=../../mocks/services/mock_query_device.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services QueryDeviceService
type QueryDeviceService interface {
	GenerateUUID() string
	GenerateRefreshToken() string
	GetOldDeviceByUserUUIDAndIpAndUserAgent(string, string, string) repositories.Device
	GetNewDeviceByUserUUIDAndIpAndUserAgent(string, string, string) (repositories.Device, error)
	GetNewRefreshDetailsByDevice(repositories.Device) (repositories.Device, error)
}

//go:generate mockgen -destination=../../mocks/services/mock_mutable_device.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services MutableDeviceService
type MutableDeviceService interface{}

//go:generate mockgen -destination=../../mocks/services/mock_device.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services DeviceService
type DeviceService interface {
	QueryDeviceService
	MutableDeviceService
}
