package repositories

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func Test_NewGormDevice(t *testing.T) {
	gormDevice := NewGormDevice()

	assert.NotNil(t, gormDevice)

	assert.Equal(t, gormDevice.UUID, GormDeviceUUIDDefault)
	assert.Equal(t, gormDevice.UserUUID, GormDeviceUserUUIDDefault)
	assert.Equal(t, gormDevice.UserAgent, GormDeviceUserAgentDefault)
	assert.Equal(t, gormDevice.Ip, GormDeviceIpDefault)
	assert.Equal(t, gormDevice.RefreshToken, GormDeviceRefreshTokenDefault)
	assert.Equal(t, gormDevice.ExpiredAt, GormDeviceExpiredAtDefault)
	assert.Equal(t, gormDevice.CreatedAt, GormDeviceCreatedAtDefault)
	assert.Equal(t, gormDevice.UpdatedAt, GormDeviceUpdatedAtDefault)
}

func Test_NewGormDeviceByDevice(t *testing.T) {
	gormDeviceTemp := NewGormDevice()
	gormDevice := NewGormDeviceByDevice(gormDeviceTemp)

	assert.NotNil(t, gormDevice)

	assert.Equal(t, gormDevice.UUID, GormDeviceUUIDDefault)
	assert.Equal(t, gormDevice.UserUUID, GormDeviceUserUUIDDefault)
	assert.Equal(t, gormDevice.UserAgent, GormDeviceUserAgentDefault)
	assert.Equal(t, gormDevice.Ip, GormDeviceIpDefault)
	assert.Equal(t, gormDevice.RefreshToken, GormDeviceRefreshTokenDefault)
	assert.Equal(t, gormDevice.ExpiredAt, GormDeviceExpiredAtDefault)
	assert.Equal(t, gormDevice.CreatedAt, GormDeviceCreatedAtDefault)
	assert.Equal(t, gormDevice.UpdatedAt, GormDeviceUpdatedAtDefault)
}

func TestGormDeviceByDevice_GenerateAccessToken(t *testing.T) {
	gormDevice := NewGormDevice()
	config := configs.NewBaseConfig()

	gormDevice.SetUUID("1")
	gormDevice.SetUserUUID("2")
	gormDevice.SetUserAgent("3")

	accessToken, err := gormDevice.GenerateAccessToken(config)

	assert.Nil(t, err)
	assert.NotNil(t, accessToken)

	assert.Equal(t, accessToken.DeviceUUID, gormDevice.GetUUID())
	assert.Equal(t, accessToken.Subject, gormDevice.GetUserUUID())
	assert.Equal(t, accessToken.DeviceUserAgent, gormDevice.GetUserAgent())
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

func TestGormDeviceByDevice_GetUserAgent(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.UserAgent = "1"

	assert.Equal(t, gormDevice.GetUserAgent(), "1")
}

func TestGormDeviceByDevice_GetIp(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.Ip = "1"

	assert.Equal(t, gormDevice.GetIp(), "1")
}

func TestGormDeviceByDevice_GetRefreshToken(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.RefreshToken = "1"

	assert.Equal(t, gormDevice.GetRefreshToken(), "1")
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

func TestGormDeviceByDevice_SetUserAgent(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetUserAgent("1")

	assert.Equal(t, gormDevice.UserAgent, "1")
}

func TestGormDeviceByDevice_SetIp(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetIp("1")

	assert.Equal(t, gormDevice.Ip, "1")
}

func TestGormDeviceByDevice_SetRefreshToken(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetRefreshToken("1")

	assert.Equal(t, gormDevice.RefreshToken, "1")
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
	assert.NotNil(t, db)

	gormDeviceRepository, err := NewGormDeviceRepository(db)

	assert.Nil(t, err)
	assert.NotNil(t, gormDeviceRepository)

	stmt := &gorm.Statement{DB: db}
	err = stmt.Parse(&GormDevice{})

	assert.Nil(t, err)

	var count int

	db.
		Raw(
			"SELECT count(*) FROM sqlite_master WHERE type = 'table' AND name = ?",
			stmt.Table,
		).
		Scan(&count)

	assert.Equal(t, count, 1)
}

func TestGormDeviceRepository_CreateDevice_And_GetDevicesByUserUUID(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormDeviceRepository, _ := NewGormDeviceRepository(db)

	gormDevice := NewGormDevice()

	gormDevice.SetUUID("1")
	gormDevice.SetUserUUID("2")
	gormDevice.SetUserAgent("3")
	gormDevice.SetIp("4")
	gormDevice.SetRefreshToken("5")
	gormDevice.SetExpiredAt(6)
	gormDevice.SetCreatedAt(7)
	gormDevice.SetUpdatedAt(8)

	err := gormDeviceRepository.CreateDevice(gormDevice)

	assert.Nil(t, err)

	gormDevices := gormDeviceRepository.GetDevicesByUserUUID(gormDevice.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)

	assert.Equal(t, gormDevices[0].GetUUID(), gormDevice.GetUUID())
	assert.Equal(t, gormDevices[0].GetUserUUID(), gormDevice.GetUserUUID())
	assert.Equal(t, gormDevices[0].GetUserAgent(), gormDevice.GetUserAgent())
	assert.Equal(t, gormDevices[0].GetIp(), gormDevice.GetIp())
	assert.Equal(t, gormDevices[0].GetExpiredAt(), gormDevice.GetExpiredAt())
	assert.Equal(t, gormDevices[0].GetCreatedAt(), gormDevice.GetCreatedAt())
	assert.Equal(t, gormDevices[0].GetUpdatedAt(), gormDevice.GetUpdatedAt())
}

func TestGormDeviceRepository_CreateDevice_And_GetDeviceByUserUUIDAndIpAndUserAgent(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormDeviceRepository, _ := NewGormDeviceRepository(db)

	gormDevice := NewGormDevice()

	gormDevice.SetUUID("1")
	gormDevice.SetUserUUID("2")
	gormDevice.SetUserAgent("3")
	gormDevice.SetIp("4")
	gormDevice.SetRefreshToken("5")
	gormDevice.SetExpiredAt(6)
	gormDevice.SetCreatedAt(7)
	gormDevice.SetUpdatedAt(8)

	err := gormDeviceRepository.CreateDevice(gormDevice)

	assert.Nil(t, err)

	gormDeviceResult := gormDeviceRepository.
		GetDeviceByUserUUIDAndIpAndUserAgent("0", "0", "0")

	assert.Nil(t, gormDeviceResult)

	gormDeviceResult = gormDeviceRepository.GetDeviceByUserUUIDAndIpAndUserAgent(
		gormDevice.GetUserUUID(),
		gormDevice.GetIp(),
		gormDevice.GetUserAgent(),
	)

	assert.NotNil(t, gormDeviceResult)

	assert.Equal(t, gormDeviceResult.GetUUID(), gormDevice.GetUUID())
	assert.Equal(t, gormDeviceResult.GetUserUUID(), gormDevice.GetUserUUID())
	assert.Equal(t, gormDeviceResult.GetUserAgent(), gormDevice.GetUserAgent())
	assert.Equal(t, gormDeviceResult.GetIp(), gormDevice.GetIp())
	assert.Equal(t, gormDeviceResult.GetExpiredAt(), gormDevice.GetExpiredAt())
	assert.Equal(t, gormDeviceResult.GetCreatedAt(), gormDevice.GetCreatedAt())
	assert.Equal(t, gormDeviceResult.GetUpdatedAt(), gormDevice.GetUpdatedAt())
}

func TestGormDeviceRepository_CreateDevice_And_GetDeviceByRefreshToken(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormDeviceRepository, _ := NewGormDeviceRepository(db)

	gormDevice := NewGormDevice()

	gormDevice.SetUUID("1")
	gormDevice.SetUserUUID("2")
	gormDevice.SetUserAgent("3")
	gormDevice.SetIp("4")
	gormDevice.SetRefreshToken("5")
	gormDevice.SetExpiredAt(6)
	gormDevice.SetCreatedAt(7)
	gormDevice.SetUpdatedAt(8)

	err := gormDeviceRepository.CreateDevice(gormDevice)

	assert.Nil(t, err)

	gormDeviceResult := gormDeviceRepository.GetDeviceByRefreshToken("0")

	assert.Nil(t, gormDeviceResult)

	gormDeviceResult = gormDeviceRepository.GetDeviceByRefreshToken(
		gormDevice.GetRefreshToken(),
	)

	assert.NotNil(t, gormDeviceResult)

	assert.Equal(t, gormDeviceResult.GetUUID(), gormDevice.GetUUID())
	assert.Equal(t, gormDeviceResult.GetUserUUID(), gormDevice.GetUserUUID())
	assert.Equal(t, gormDeviceResult.GetUserAgent(), gormDevice.GetUserAgent())
	assert.Equal(t, gormDeviceResult.GetIp(), gormDevice.GetIp())
	assert.Equal(t, gormDeviceResult.GetExpiredAt(), gormDevice.GetExpiredAt())
	assert.Equal(t, gormDeviceResult.GetCreatedAt(), gormDevice.GetCreatedAt())
	assert.Equal(t, gormDeviceResult.GetUpdatedAt(), gormDevice.GetUpdatedAt())
}

func TestGormDeviceRepository_CreateDevice_And_UpdateDevice(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

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
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

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

func TestGormDeviceRepository_CreateDevice_And_DeleteDeviceByUUIDAndUserUUID(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormDeviceRepository, _ := NewGormDeviceRepository(db)

	gormDevice := NewGormDevice()

	gormDevice.SetUUID("1")
	gormDevice.SetUserUUID("2")

	err := gormDeviceRepository.CreateDevice(gormDevice)

	assert.Nil(t, err)

	gormDevices := gormDeviceRepository.GetDevicesByUserUUID(gormDevice.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)

	err = gormDeviceRepository.
		DeleteDeviceByUUIDAndUserUUID(gormDevice.GetUUID(), gormDevice.GetUserUUID())

	assert.Nil(t, err)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDevice.GetUserUUID())

	assert.Equal(t, len(gormDevices), 0)
}

func TestGormDeviceRepository_CreateDevice_And_DeleteAllDevicesByUserUUID(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormDeviceRepository, _ := NewGormDeviceRepository(db)

	gormDevice := NewGormDevice()

	gormDevice.SetUUID("1")
	gormDevice.SetUserUUID("2")

	err := gormDeviceRepository.CreateDevice(gormDevice)

	assert.Nil(t, err)

	gormDevice.SetUUID("3")
	gormDevice.SetUserUUID("2")

	err = gormDeviceRepository.CreateDevice(gormDevice)

	assert.Nil(t, err)

	gormDevices := gormDeviceRepository.GetDevicesByUserUUID(gormDevice.GetUserUUID())

	assert.Equal(t, len(gormDevices), 2)

	err = gormDeviceRepository.DeleteAllDevicesByUserUUID(gormDevice.GetUserUUID())

	assert.Nil(t, err)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDevice.GetUserUUID())

	assert.Equal(t, len(gormDevices), 0)
}
