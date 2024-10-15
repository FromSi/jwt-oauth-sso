package repositories

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func Test_NewGormDevice(t *testing.T) {
	gormDevice := NewGormDevice()

	assert.Equal(t, gormDevice.UUID, GormDeviceUUIDDefault)
	assert.Equal(t, gormDevice.UserUUID, GormDeviceUserUUIDDefault)
	assert.Equal(t, gormDevice.Agent, GormDeviceAgentDefault)
	assert.Equal(t, gormDevice.Ip, GormDeviceIpDefault)
	assert.Equal(t, gormDevice.ExpiredAt, GormDeviceExpiredAtDefault)
	assert.Equal(t, gormDevice.CreatedAt, GormDeviceCreatedAtDefault)
	assert.Equal(t, gormDevice.UpdatedAt, GormDeviceUpdatedAtDefault)
}

func Test_NewGormDeviceByDevice(t *testing.T) {
	gormDeviceTemp := NewGormDevice()
	gormDevice := NewGormDeviceByDevice(gormDeviceTemp)

	assert.Equal(t, gormDevice.UUID, GormDeviceUUIDDefault)
	assert.Equal(t, gormDevice.UserUUID, GormDeviceUserUUIDDefault)
	assert.Equal(t, gormDevice.Agent, GormDeviceAgentDefault)
	assert.Equal(t, gormDevice.Ip, GormDeviceIpDefault)
	assert.Equal(t, gormDevice.ExpiredAt, GormDeviceExpiredAtDefault)
	assert.Equal(t, gormDevice.CreatedAt, GormDeviceCreatedAtDefault)
	assert.Equal(t, gormDevice.UpdatedAt, GormDeviceUpdatedAtDefault)
}

func TestGormDeviceByDevice_GetUUID(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.UUID = "1"

	assert.Equal(t, gormDevice.GetUUID(), "1")
}

func TestGormDeviceByDevice_GetUserUUID(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.UserUUID = "1"

	assert.Equal(t, gormDevice.GetUserUUID(), "1")
}

func TestGormDeviceByDevice_GetAgent(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.Agent = "1"

	assert.Equal(t, gormDevice.GetAgent(), "1")
}

func TestGormDeviceByDevice_GetIp(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.Ip = "1"

	assert.Equal(t, gormDevice.GetIp(), "1")
}

func TestGormDeviceByDevice_GetExpiredAt(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.ExpiredAt = 1

	assert.Equal(t, gormDevice.GetExpiredAt(), 1)
}

func TestGormDeviceByDevice_GetCreatedAt(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.CreatedAt = 1

	assert.Equal(t, gormDevice.GetCreatedAt(), 1)
}

func TestGormDeviceByDevice_GetUpdatedAt(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.UpdatedAt = 1

	assert.Equal(t, gormDevice.GetUpdatedAt(), 1)
}

func TestGormDeviceByDevice_SetUUID(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetUUID("1")

	assert.Equal(t, gormDevice.UUID, "1")
}

func TestGormDeviceByDevice_SetUserUUID(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetUserUUID("1")

	assert.Equal(t, gormDevice.UserUUID, "1")
}

func TestGormDeviceByDevice_SetAgent(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetAgent("1")

	assert.Equal(t, gormDevice.Agent, "1")
}

func TestGormDeviceByDevice_SetIp(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetIp("1")

	assert.Equal(t, gormDevice.Ip, "1")
}

func TestGormDeviceByDevice_SetExpiredAt(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetExpiredAt(1)

	assert.Equal(t, gormDevice.ExpiredAt, 1)
}

func TestGormDeviceByDevice_SetCreatedAt(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetCreatedAt(1)

	assert.Equal(t, gormDevice.CreatedAt, 1)
}

func TestGormDeviceByDevice_SetUpdatedAt(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetUpdatedAt(1)

	assert.Equal(t, gormDevice.UpdatedAt, 1)
}

func Test_NewGormDeviceRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	assert.Nil(t, err)

	_, err = NewGormDeviceRepository(db)

	assert.Nil(t, err)

	stmt := &gorm.Statement{DB: db}
	err = stmt.Parse(&GormDevice{})

	assert.Nil(t, err)

	var count int

	db.Raw("SELECT count(*) FROM sqlite_master WHERE type = 'table' AND name = ?", stmt.Table).Scan(&count)

	assert.Equal(t, count, 1)
}

func TestGormDeviceRepository_CreateDevice_And_GetDevicesByUserUUID(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	gormDeviceRepository, _ := NewGormDeviceRepository(db)

	gormDevice := NewGormDevice()

	gormDevice.SetUUID("1")
	gormDevice.SetUserUUID("2")
	gormDevice.SetAgent("3")
	gormDevice.SetIp("4")
	gormDevice.SetExpiredAt(5)
	gormDevice.SetCreatedAt(6)
	gormDevice.SetUpdatedAt(7)

	err := gormDeviceRepository.CreateDevice(gormDevice)

	assert.Nil(t, err)

	gormDevices := gormDeviceRepository.GetDevicesByUserUUID(gormDevice.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)

	assert.Equal(t, gormDevices[0].GetUUID(), gormDevice.GetUUID())
	assert.Equal(t, gormDevices[0].GetUserUUID(), gormDevice.GetUserUUID())
	assert.Equal(t, gormDevices[0].GetAgent(), gormDevice.GetAgent())
	assert.Equal(t, gormDevices[0].GetIp(), gormDevice.GetIp())
	assert.Equal(t, gormDevices[0].GetExpiredAt(), gormDevice.GetExpiredAt())
	assert.Equal(t, gormDevices[0].GetCreatedAt(), gormDevice.GetCreatedAt())
	assert.Equal(t, gormDevices[0].GetUpdatedAt(), gormDevice.GetUpdatedAt())
}

func TestGormDeviceRepository_CreateDevice_And_UpdateDevice(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	gormDeviceRepository, _ := NewGormDeviceRepository(db)

	gormDevice := NewGormDevice()

	gormDevice.SetUUID("1")
	gormDevice.SetUserUUID("2")
	gormDevice.SetUpdatedAt(3)

	err := gormDeviceRepository.CreateDevice(gormDevice)

	assert.Nil(t, err)

	gormDevices := gormDeviceRepository.GetDevicesByUserUUID(gormDevice.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)
	assert.Equal(t, gormDevices[0].GetUserUUID(), gormDevice.GetUserUUID())
	assert.Equal(t, gormDevices[0].GetUpdatedAt(), gormDevice.GetUpdatedAt())

	gormDevice.SetUserUUID("4")
	gormDevice.SetUpdatedAt(5)

	err = gormDeviceRepository.UpdateDevice(gormDevice)

	assert.Nil(t, err)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDevice.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)
	assert.Equal(t, gormDevices[0].GetUserUUID(), gormDevice.GetUserUUID())
	assert.Equal(t, gormDevices[0].GetUpdatedAt(), gormDevice.GetUpdatedAt())
}

func TestGormDeviceRepository_CreateDevice_And_DeleteDeviceByUUID(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	gormDeviceRepository, _ := NewGormDeviceRepository(db)

	gormDevice := NewGormDevice()

	gormDevice.SetUUID("1")
	gormDevice.SetUserUUID("2")

	err := gormDeviceRepository.CreateDevice(gormDevice)

	assert.Nil(t, err)

	gormDevices := gormDeviceRepository.GetDevicesByUserUUID(gormDevice.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)

	err = gormDeviceRepository.DeleteDeviceByUUID(gormDevice.GetUUID())

	assert.Nil(t, err)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDevice.GetUserUUID())

	assert.Equal(t, len(gormDevices), 0)
}
