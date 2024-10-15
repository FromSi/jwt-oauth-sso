package repositories

import (
	"gorm.io/gorm"
)

const (
	GormDeviceUUIDDefault      = ""
	GormDeviceUserUUIDDefault  = ""
	GormDeviceAgentDefault     = ""
	GormDeviceIpDefault        = ""
	GormDeviceExpiredAtDefault = 0
	GormDeviceCreatedAtDefault = 0
	GormDeviceUpdatedAtDefault = 0
)

type GormDevice struct {
	UUID      string `gorm:"unique;not null"`
	UserUUID  string `gorm:"not null"`
	Agent     string `gorm:"not null"`
	Ip        string `gorm:"not null"`
	ExpiredAt int    `gorm:"not null"`
	CreatedAt int    `gorm:"not null"`
	UpdatedAt int    `gorm:"not null"`
}

func NewGormDevice() *GormDevice {
	return &GormDevice{
		UUID:      GormDeviceUUIDDefault,
		UserUUID:  GormDeviceUserUUIDDefault,
		Agent:     GormDeviceAgentDefault,
		Ip:        GormDeviceIpDefault,
		ExpiredAt: GormDeviceExpiredAtDefault,
		CreatedAt: GormDeviceCreatedAtDefault,
		UpdatedAt: GormDeviceUpdatedAtDefault,
	}
}

func NewGormDeviceByDevice(device Device) *GormDevice {
	return &GormDevice{
		UUID:      device.GetUUID(),
		UserUUID:  device.GetUserUUID(),
		Agent:     device.GetAgent(),
		Ip:        device.GetIp(),
		ExpiredAt: device.GetExpiredAt(),
		CreatedAt: device.GetCreatedAt(),
		UpdatedAt: device.GetUpdatedAt(),
	}
}

func (receiver GormDevice) GetUUID() string {
	return receiver.UUID
}

func (receiver GormDevice) GetUserUUID() string {
	return receiver.UserUUID
}

func (receiver GormDevice) GetAgent() string {
	return receiver.Agent
}

func (receiver GormDevice) GetIp() string {
	return receiver.Ip
}

func (receiver GormDevice) GetExpiredAt() int {
	return receiver.ExpiredAt
}

func (receiver GormDevice) GetCreatedAt() int {
	return receiver.CreatedAt
}

func (receiver GormDevice) GetUpdatedAt() int {
	return receiver.UpdatedAt
}

func (receiver *GormDevice) SetUUID(value string) {
	receiver.UUID = value
}

func (receiver *GormDevice) SetUserUUID(value string) {
	receiver.UserUUID = value
}

func (receiver *GormDevice) SetAgent(value string) {
	receiver.Agent = value
}

func (receiver *GormDevice) SetIp(value string) {
	receiver.Ip = value
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

func (receiver GormDeviceRepository) GetDevicesByUserUUID(userUUID string) []Device {
	var gormDevices []GormDevice

	receiver.db.Model(&GormDevice{UserUUID: userUUID}).Find(&gormDevices)

	devices := make([]Device, len(gormDevices))

	for i, gormDevice := range gormDevices {
		devices[i] = &gormDevice
	}

	return devices
}

func (receiver GormDeviceRepository) CreateDevice(device Device) error {
	gormDevice := NewGormDeviceByDevice(device)

	return receiver.db.Model(&GormDevice{}).Create(NewGormDeviceByDevice(gormDevice)).Error
}

func (receiver GormDeviceRepository) UpdateDevice(device Device) error {
	gormDevice := NewGormDeviceByDevice(device)

	return receiver.db.Model(&GormDevice{}).Where(&GormDevice{UUID: device.GetUUID()}).UpdateColumns(NewGormDeviceByDevice(gormDevice)).Error
}

func (receiver GormDeviceRepository) DeleteDeviceByUUID(uuid string) error {
	return receiver.db.Delete(&GormDevice{}, &GormDevice{UUID: uuid}).Error
}
