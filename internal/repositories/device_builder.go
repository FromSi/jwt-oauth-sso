package repositories

//go:generate mockgen -destination=../../mocks/repositories/mock_device_builder.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories DeviceBuilder
type DeviceBuilder interface {
	New() DeviceBuilder
	NewFromDevice(Device) DeviceBuilder
	Build() (Device, error)
	BuildToGorm() (*GormDevice, error)
	SetUUID(string) DeviceBuilder
	SetUserUUID(string) DeviceBuilder
	SetUserAgent(string) DeviceBuilder
	SetIp(string) DeviceBuilder
	SetRefreshToken(string) DeviceBuilder
	SetIssuedAt(int) DeviceBuilder
	SetExpiresAt(int) DeviceBuilder
	SetCreatedAt(int) DeviceBuilder
	SetUpdatedAt(int) DeviceBuilder
}
