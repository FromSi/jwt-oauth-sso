package repositories

import (
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
	"gorm.io/gorm"
)

type GormDevice struct {
	UUID               string `gorm:"unique;not null"`
	UserUUID           string `gorm:"not null"`
	UserAgent          string `gorm:"not null"`
	Ip                 string `gorm:"not null"`
	RefreshToken       string `gorm:"not null"`
	IssuedAt           int    `gorm:"not null"`
	ExpiresAt          int    `gorm:"not null"`
	CreatedAt          int    `gorm:"not null"`
	UpdatedAt          int    `gorm:"not null"`
	accessTokenBuilder tokens.AccessTokenBuilder
}

func (receiver *GormDevice) TableName() string {
	return "devices"
}

func (receiver *GormDevice) GenerateAccessToken() (tokens.AccessToken, error) {
	return receiver.
		accessTokenBuilder.
		New().
		SetSubject(receiver.UserUUID).
		SetDeviceUUID(receiver.UUID).
		SetDeviceUserAgent(receiver.UserAgent).
		SetIssuedAt(receiver.IssuedAt).
		SetExpirationTime(receiver.ExpiresAt).
		Build()
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

func (receiver *GormDevice) GetIssuedAt() int {
	return receiver.IssuedAt
}

func (receiver *GormDevice) GetExpiresAt() int {
	return receiver.ExpiresAt
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

func (receiver *GormDevice) SetIssuedAt(value int) {
	receiver.IssuedAt = value
}

func (receiver *GormDevice) SetExpiresAt(value int) {
	receiver.ExpiresAt = value
}

func (receiver *GormDevice) SetCreatedAt(value int) {
	receiver.CreatedAt = value
}

func (receiver *GormDevice) SetUpdatedAt(value int) {
	receiver.UpdatedAt = value
}

type GormDeviceRepository struct {
	db            *gorm.DB
	deviceBuilder DeviceBuilder
}

func NewGormDeviceRepository(
	db *gorm.DB,
	deviceBuilder DeviceBuilder,
) (*GormDeviceRepository, error) {
	err := db.AutoMigrate(&GormDevice{})

	if err != nil {
		return nil, err
	}

	return &GormDeviceRepository{
		db:            db,
		deviceBuilder: deviceBuilder,
	}, nil
}

func (receiver *GormDeviceRepository) GetDevicesByUserUUID(userUUID string) []Device {
	var gormDevices []GormDevice

	receiver.
		db.
		Model(&GormDevice{}).
		Find(&gormDevices, &GormDevice{UserUUID: userUUID})

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
	gormDevice, err := receiver.
		deviceBuilder.
		NewFromDevice(device).
		BuildToGorm()

	if err != nil {
		return err
	}

	return receiver.
		db.
		Model(&GormDevice{}).
		Create(gormDevice).
		Error
}

func (receiver *GormDeviceRepository) UpdateDevice(device Device) error {
	gormDevice, err := receiver.
		deviceBuilder.
		NewFromDevice(device).
		BuildToGorm()

	if err != nil {
		return err
	}

	return receiver.
		db.
		Model(&GormDevice{}).
		Where(&GormDevice{UUID: device.GetUUID()}).
		UpdateColumns(gormDevice).
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
