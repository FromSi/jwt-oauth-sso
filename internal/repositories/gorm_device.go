package repositories

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
	"gorm.io/gorm"
	"time"
)

const (
	GormDeviceUUIDDefault         = ""
	GormDeviceUserUUIDDefault     = ""
	GormDeviceUserAgentDefault    = ""
	GormDeviceIpDefault           = ""
	GormDeviceRefreshTokenDefault = ""
	GormDeviceExpiredAtDefault    = 0
	GormDeviceCreatedAtDefault    = 0
	GormDeviceUpdatedAtDefault    = 0
)

type GormDevice struct {
	UUID         string `gorm:"unique;not null"`
	UserUUID     string `gorm:"not null"`
	UserAgent    string `gorm:"not null"`
	Ip           string `gorm:"not null"`
	RefreshToken string `gorm:"not null"`
	ExpiredAt    int    `gorm:"not null"`
	CreatedAt    int    `gorm:"not null"`
	UpdatedAt    int    `gorm:"not null"`
}

func NewGormDevice() *GormDevice {
	return &GormDevice{
		UUID:         GormDeviceUUIDDefault,
		UserUUID:     GormDeviceUserUUIDDefault,
		UserAgent:    GormDeviceUserAgentDefault,
		Ip:           GormDeviceIpDefault,
		RefreshToken: GormDeviceRefreshTokenDefault,
		ExpiredAt:    GormDeviceExpiredAtDefault,
		CreatedAt:    GormDeviceCreatedAtDefault,
		UpdatedAt:    GormDeviceUpdatedAtDefault,
	}
}

func NewGormDeviceByDevice(device Device) *GormDevice {
	return &GormDevice{
		UUID:         device.GetUUID(),
		UserUUID:     device.GetUserUUID(),
		UserAgent:    device.GetUserAgent(),
		Ip:           device.GetIp(),
		RefreshToken: device.GetRefreshToken(),
		ExpiredAt:    device.GetExpiredAt(),
		CreatedAt:    device.GetCreatedAt(),
		UpdatedAt:    device.GetUpdatedAt(),
	}
}

func (receiver *GormDevice) GenerateAccessToken(
	config configs.TokenConfig,
) (*tokens.AccessToken, error) {
	return tokens.NewAccessToken(
		config,
		receiver.UserUUID,
		receiver.UUID,
		receiver.UserAgent,
		time.Now(),
	)
}

func (receiver *GormDevice) GetUUID() string {
	return receiver.UUID
}

func (receiver *GormDevice) GetUserUUID() string {
	return receiver.UserUUID
}

func (receiver *GormDevice) GetUserAgent() string {
	return receiver.UserAgent
}

func (receiver *GormDevice) GetIp() string {
	return receiver.Ip
}

func (receiver *GormDevice) GetRefreshToken() string {
	return receiver.RefreshToken
}

func (receiver *GormDevice) GetExpiredAt() int {
	return receiver.ExpiredAt
}

func (receiver *GormDevice) GetCreatedAt() int {
	return receiver.CreatedAt
}

func (receiver *GormDevice) GetUpdatedAt() int {
	return receiver.UpdatedAt
}

func (receiver *GormDevice) SetUUID(value string) {
	receiver.UUID = value
}

func (receiver *GormDevice) SetUserUUID(value string) {
	receiver.UserUUID = value
}

func (receiver *GormDevice) SetUserAgent(value string) {
	receiver.UserAgent = value
}

func (receiver *GormDevice) SetIp(value string) {
	receiver.Ip = value
}

func (receiver *GormDevice) SetRefreshToken(value string) {
	receiver.RefreshToken = value
}

func (receiver *GormDevice) SetExpiredAt(value int) {
	receiver.ExpiredAt = value
}

func (receiver *GormDevice) SetCreatedAt(value int) {
	receiver.CreatedAt = value
}

func (receiver *GormDevice) SetUpdatedAt(value int) {
	receiver.UpdatedAt = value
}

type GormDeviceRepository struct {
	db *gorm.DB
}

func NewGormDeviceRepository(db *gorm.DB) (*GormDeviceRepository, error) {
	err := db.AutoMigrate(&GormDevice{})

	if err != nil {
		return nil, err
	}

	return &GormDeviceRepository{db: db}, nil
}

func (receiver *GormDeviceRepository) GetDevicesByUserUUID(userUUID string) []Device {
	var gormDevices []GormDevice

	receiver.
		db.
		Model(&GormDevice{UserUUID: userUUID}).
		Find(&gormDevices)

	devices := make([]Device, len(gormDevices))

	for i, gormDevice := range gormDevices {
		devices[i] = &gormDevice
	}

	return devices
}

func (receiver *GormDeviceRepository) GetDeviceByUserUUIDAndIpAndUserAgent(
	userUUID string,
	ip string,
	agent string,
) Device {
	var gormDevice GormDevice

	result := receiver.
		db.
		Model(&GormDevice{}).
		First(&gormDevice, &GormDevice{UserUUID: userUUID, Ip: ip, UserAgent: agent})

	if result.RowsAffected == 0 {
		return nil
	}

	return &gormDevice
}

func (receiver *GormDeviceRepository) GetDeviceByRefreshToken(refreshToken string) Device {
	var gormDevice GormDevice

	result := receiver.
		db.
		Model(&GormDevice{}).
		First(&gormDevice, &GormDevice{RefreshToken: refreshToken})

	if result.RowsAffected == 0 {
		return nil
	}

	return &gormDevice
}

func (receiver *GormDeviceRepository) CreateDevice(device Device) error {
	gormDevice := NewGormDeviceByDevice(device)

	return receiver.
		db.
		Model(&GormDevice{}).
		Create(NewGormDeviceByDevice(gormDevice)).
		Error
}

func (receiver *GormDeviceRepository) UpdateDevice(device Device) error {
	gormDevice := NewGormDeviceByDevice(device)

	return receiver.
		db.
		Model(&GormDevice{}).
		Where(&GormDevice{UUID: device.GetUUID()}).
		UpdateColumns(NewGormDeviceByDevice(gormDevice)).
		Error
}

func (receiver *GormDeviceRepository) DeleteDeviceByUUID(uuid string) error {
	return receiver.
		db.
		Delete(&GormDevice{}, &GormDevice{UUID: uuid}).
		Error
}

func (receiver *GormDeviceRepository) DeleteDeviceByUUIDAndUserUUID(
	uuid string,
	userUUID string,
) error {
	return receiver.
		db.
		Delete(&GormDevice{}, &GormDevice{UUID: uuid, UserUUID: userUUID}).
		Error
}

func (receiver *GormDeviceRepository) DeleteAllDevicesByUserUUID(userUUID string) error {
	return receiver.
		db.
		Delete(&GormDevice{}, &GormDevice{UserUUID: userUUID}).
		Error
}
