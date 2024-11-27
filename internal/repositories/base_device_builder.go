package repositories

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
)

type BaseDeviceBuilder struct {
	device GormDevice
}

func NewBaseDeviceBuilder(accessTokenBuilder tokens.AccessTokenBuilder) *BaseDeviceBuilder {
	return &BaseDeviceBuilder{
		device: GormDevice{accessTokenBuilder: accessTokenBuilder},
	}
}

func (receiver *BaseDeviceBuilder) New() DeviceBuilder {
	return &BaseDeviceBuilder{
		device: GormDevice{accessTokenBuilder: receiver.device.accessTokenBuilder},
	}
}

func (receiver *BaseDeviceBuilder) NewFromDevice(device Device) DeviceBuilder {
	return receiver.
		New().
		SetUUID(device.GetUUID()).
		SetUserUUID(device.GetUserUUID()).
		SetUserAgent(device.GetUserAgent()).
		SetIp(device.GetIp()).
		SetRefreshToken(device.GetRefreshToken()).
		SetIssuedAt(device.GetIssuedAt()).
		SetExpiresAt(device.GetExpiresAt()).
		SetCreatedAt(device.GetCreatedAt()).
		SetUpdatedAt(device.GetUpdatedAt())
}

func (receiver *BaseDeviceBuilder) Build() (Device, error) {
	return receiver.BuildToGorm()
}

func (receiver *BaseDeviceBuilder) BuildToGorm() (*GormDevice, error) {
	if len(receiver.device.GetUUID()) == 0 {
		return nil, errors.New("uuid must not be empty")
	}

	return &receiver.device, nil
}

func (receiver *BaseDeviceBuilder) SetUUID(value string) DeviceBuilder {
	receiver.device.SetUUID(value)

	return receiver
}

func (receiver *BaseDeviceBuilder) SetUserUUID(value string) DeviceBuilder {
	receiver.device.SetUserUUID(value)

	return receiver
}

func (receiver *BaseDeviceBuilder) SetUserAgent(value string) DeviceBuilder {
	receiver.device.SetUserAgent(value)

	return receiver
}

func (receiver *BaseDeviceBuilder) SetIp(value string) DeviceBuilder {
	receiver.device.SetIp(value)

	return receiver
}

func (receiver *BaseDeviceBuilder) SetRefreshToken(value string) DeviceBuilder {
	receiver.device.SetRefreshToken(value)

	return receiver
}

func (receiver *BaseDeviceBuilder) SetIssuedAt(value int) DeviceBuilder {
	receiver.device.SetIssuedAt(value)

	return receiver
}

func (receiver *BaseDeviceBuilder) SetExpiresAt(value int) DeviceBuilder {
	receiver.device.SetExpiresAt(value)

	return receiver
}

func (receiver *BaseDeviceBuilder) SetCreatedAt(value int) DeviceBuilder {
	receiver.device.SetCreatedAt(value)

	return receiver
}

func (receiver *BaseDeviceBuilder) SetUpdatedAt(value int) DeviceBuilder {
	receiver.device.SetUpdatedAt(value)

	return receiver
}
