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
	assert.Equal(t, gormDevice.ExpiresAt, GormDeviceExpiresAtDefault)
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
	assert.Equal(t, gormDevice.ExpiresAt, GormDeviceExpiresAtDefault)
	assert.Equal(t, gormDevice.CreatedAt, GormDeviceCreatedAtDefault)
	assert.Equal(t, gormDevice.UpdatedAt, GormDeviceUpdatedAtDefault)
}

func TestGormDeviceByDevice_GenerateAccessToken(t *testing.T) {
	gormDevice := NewGormDevice()
	config := configs.NewBaseConfig()

	gormDevice.SetUUID("1")
	gormDevice.SetUserUUID("1")
	gormDevice.SetUserAgent("1")

	accessToken, err := gormDevice.GenerateAccessToken(config)

	assert.NoError(t, err)
	assert.NotNil(t, accessToken)
}

func TestGormDeviceByDevice_GetUUID(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.UUID = "1"

	assert.Equal(t, gormDevice.GetUUID(), "1")

	gormDevice.UUID = "2"

	assert.Equal(t, gormDevice.GetUUID(), "2")
}

func TestGormDeviceByDevice_GetUserUUID(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.UserUUID = "1"

	assert.Equal(t, gormDevice.GetUserUUID(), "1")

	gormDevice.UserUUID = "2"

	assert.Equal(t, gormDevice.GetUserUUID(), "2")
}

func TestGormDeviceByDevice_GetUserAgent(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.UserAgent = "1"

	assert.Equal(t, gormDevice.GetUserAgent(), "1")

	gormDevice.UserAgent = "2"

	assert.Equal(t, gormDevice.GetUserAgent(), "2")
}

func TestGormDeviceByDevice_GetIp(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.Ip = "1"

	assert.Equal(t, gormDevice.GetIp(), "1")

	gormDevice.Ip = "2"

	assert.Equal(t, gormDevice.GetIp(), "2")
}

func TestGormDeviceByDevice_GetRefreshToken(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.RefreshToken = "1"

	assert.Equal(t, gormDevice.GetRefreshToken(), "1")

	gormDevice.RefreshToken = "2"

	assert.Equal(t, gormDevice.GetRefreshToken(), "2")
}

func TestGormDeviceByDevice_GetExpiresAt(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.ExpiresAt = 1

	assert.Equal(t, gormDevice.GetExpiresAt(), 1)

	gormDevice.ExpiresAt = 2

	assert.Equal(t, gormDevice.GetExpiresAt(), 2)
}

func TestGormDeviceByDevice_GetCreatedAt(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.CreatedAt = 1

	assert.Equal(t, gormDevice.GetCreatedAt(), 1)

	gormDevice.CreatedAt = 2

	assert.Equal(t, gormDevice.GetCreatedAt(), 2)
}

func TestGormDeviceByDevice_GetUpdatedAt(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.UpdatedAt = 1

	assert.Equal(t, gormDevice.GetUpdatedAt(), 1)

	gormDevice.UpdatedAt = 2

	assert.Equal(t, gormDevice.GetUpdatedAt(), 2)
}

func TestGormDeviceByDevice_SetUUID(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetUUID("1")

	assert.Equal(t, gormDevice.UUID, "1")

	gormDevice.SetUUID("2")

	assert.Equal(t, gormDevice.UUID, "2")
}

func TestGormDeviceByDevice_SetUserUUID(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetUserUUID("1")

	assert.Equal(t, gormDevice.UserUUID, "1")

	gormDevice.SetUserUUID("2")

	assert.Equal(t, gormDevice.UserUUID, "2")
}

func TestGormDeviceByDevice_SetUserAgent(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetUserAgent("1")

	assert.Equal(t, gormDevice.UserAgent, "1")

	gormDevice.SetUserAgent("2")

	assert.Equal(t, gormDevice.UserAgent, "2")
}

func TestGormDeviceByDevice_SetIp(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetIp("1")

	assert.Equal(t, gormDevice.Ip, "1")

	gormDevice.SetIp("2")

	assert.Equal(t, gormDevice.Ip, "2")
}

func TestGormDeviceByDevice_SetRefreshToken(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetRefreshToken("1")

	assert.Equal(t, gormDevice.RefreshToken, "1")

	gormDevice.SetRefreshToken("2")

	assert.Equal(t, gormDevice.RefreshToken, "2")
}

func TestGormDeviceByDevice_SetExpiresAt(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetExpiresAt(1)

	assert.Equal(t, gormDevice.ExpiresAt, 1)

	gormDevice.SetExpiresAt(2)

	assert.Equal(t, gormDevice.ExpiresAt, 2)
}

func TestGormDeviceByDevice_SetCreatedAt(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetCreatedAt(1)

	assert.Equal(t, gormDevice.CreatedAt, 1)

	gormDevice.SetCreatedAt(2)

	assert.Equal(t, gormDevice.CreatedAt, 2)
}

func TestGormDeviceByDevice_SetUpdatedAt(t *testing.T) {
	gormDevice := NewGormDevice()

	gormDevice.SetUpdatedAt(1)

	assert.Equal(t, gormDevice.UpdatedAt, 1)

	gormDevice.SetUpdatedAt(2)

	assert.Equal(t, gormDevice.UpdatedAt, 2)
}

func Test_NewGormDeviceRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?mode=ro"), &gorm.Config{})

	assert.NoError(t, err)
	assert.NotNil(t, db)

	gormDeviceRepository, err := NewGormDeviceRepository(db)

	assert.Error(t, err)
	assert.Nil(t, gormDeviceRepository)

	db, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	assert.NoError(t, err)
	assert.NotNil(t, db)

	stmt := &gorm.Statement{DB: db}

	var tableCount int

	db.
		Raw(
			"SELECT count(*) FROM sqlite_master WHERE type = 'table' AND name = ?",
			stmt.Table,
		).
		Scan(&tableCount)

	assert.Equal(t, tableCount, 0)

	gormDeviceRepository, err = NewGormDeviceRepository(db)

	assert.NoError(t, err)
	assert.NotNil(t, gormDeviceRepository)

	err = stmt.Parse(&GormDevice{})

	assert.NoError(t, err)

	db.
		Raw(
			"SELECT count(*) FROM sqlite_master WHERE type = 'table' AND name = ?",
			(&GormDevice{}).TableName(),
		).
		Scan(&tableCount)

	assert.Equal(t, tableCount, 1)
}

func TestGormDeviceRepository_CreateDevice_And_GetDeviceByUserUUIDAndIpAndUserAgent(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormDeviceRepository, _ := NewGormDeviceRepository(db)

	gormDevice := NewGormDevice()

	gormDevice.SetUserUUID("1")
	gormDevice.SetUserAgent("1")
	gormDevice.SetIp("1")

	err := gormDeviceRepository.CreateDevice(gormDevice)

	assert.NoError(t, err)

	gormDeviceResult := gormDeviceRepository.
		GetDeviceByUserUUIDAndIpAndUserAgent("0", "0", "0")

	assert.Nil(t, gormDeviceResult)

	gormDeviceResult = gormDeviceRepository.GetDeviceByUserUUIDAndIpAndUserAgent(
		gormDevice.GetUserUUID(),
		gormDevice.GetIp(),
		gormDevice.GetUserAgent(),
	)

	assert.NotNil(t, gormDeviceResult)

	assert.Equal(t, gormDeviceResult.GetUserUUID(), gormDevice.GetUserUUID())
	assert.Equal(t, gormDeviceResult.GetUserAgent(), gormDevice.GetUserAgent())
	assert.Equal(t, gormDeviceResult.GetIp(), gormDevice.GetIp())
}

func TestGormDeviceRepository_CreateDevice_And_GetDeviceByRefreshToken(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormDeviceRepository, _ := NewGormDeviceRepository(db)

	gormDevice := NewGormDevice()

	gormDevice.SetRefreshToken("1")

	err := gormDeviceRepository.CreateDevice(gormDevice)

	assert.NoError(t, err)

	gormDeviceResult := gormDeviceRepository.GetDeviceByRefreshToken("0")

	assert.Nil(t, gormDeviceResult)

	gormDeviceResult = gormDeviceRepository.GetDeviceByRefreshToken(
		gormDevice.GetRefreshToken(),
	)

	assert.NotNil(t, gormDeviceResult)

	assert.Equal(t, gormDeviceResult.GetRefreshToken(), gormDevice.GetRefreshToken())
}

func TestGormDeviceRepository_CreateDevice_And_GetDevicesByUserUUID_And_UpdateDevice(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormDeviceRepository, _ := NewGormDeviceRepository(db)

	gormDeviceOne := NewGormDevice()

	gormDeviceOne.SetUUID("1")
	gormDeviceOne.SetUserUUID("1")
	gormDeviceOne.SetUpdatedAt(1)

	err := gormDeviceRepository.CreateDevice(gormDeviceOne)

	assert.NoError(t, err)

	gormDeviceTwo := NewGormDevice()

	gormDeviceTwo.SetUUID("2")
	gormDeviceTwo.SetUserUUID("2")
	gormDeviceTwo.SetUpdatedAt(2)

	err = gormDeviceRepository.CreateDevice(gormDeviceTwo)

	assert.NoError(t, err)

	gormDevices := gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)
	assert.Equal(t, gormDevices[0].GetUserUUID(), gormDeviceOne.GetUserUUID())
	assert.Equal(t, gormDevices[0].GetUpdatedAt(), gormDeviceOne.GetUpdatedAt())

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceTwo.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)
	assert.Equal(t, gormDevices[0].GetUserUUID(), gormDeviceTwo.GetUserUUID())
	assert.Equal(t, gormDevices[0].GetUpdatedAt(), gormDeviceTwo.GetUpdatedAt())

	gormDeviceOne.SetUserUUID("3")
	gormDeviceOne.SetUpdatedAt(3)

	err = gormDeviceRepository.UpdateDevice(gormDeviceOne)

	assert.NoError(t, err)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)
	assert.Equal(t, gormDevices[0].GetUserUUID(), gormDeviceOne.GetUserUUID())
	assert.Equal(t, gormDevices[0].GetUpdatedAt(), gormDeviceOne.GetUpdatedAt())

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceTwo.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)
	assert.Equal(t, gormDevices[0].GetUserUUID(), gormDeviceTwo.GetUserUUID())
	assert.Equal(t, gormDevices[0].GetUpdatedAt(), gormDeviceTwo.GetUpdatedAt())
}

func TestGormDeviceRepository_CreateDevice_And_GetDevicesByUserUUID_And_DeleteDeviceByUUID(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormDeviceRepository, _ := NewGormDeviceRepository(db)

	gormDeviceOne := NewGormDevice()

	gormDeviceOne.SetUUID("1")
	gormDeviceOne.SetUserUUID("1")

	err := gormDeviceRepository.CreateDevice(gormDeviceOne)

	assert.NoError(t, err)

	gormDeviceTwo := NewGormDevice()

	gormDeviceTwo.SetUUID("2")
	gormDeviceTwo.SetUserUUID("2")

	err = gormDeviceRepository.CreateDevice(gormDeviceTwo)

	assert.NoError(t, err)

	gormDevices := gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceTwo.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)

	err = gormDeviceRepository.DeleteDeviceByUUID(gormDeviceOne.GetUUID())

	assert.NoError(t, err)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, len(gormDevices), 0)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceTwo.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)
}

func TestGormDeviceRepository_CreateDevice_And_DeleteDeviceByUUIDAndUserUUID(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormDeviceRepository, _ := NewGormDeviceRepository(db)

	gormDeviceOne := NewGormDevice()

	gormDeviceOne.SetUUID("1")
	gormDeviceOne.SetUserUUID("1")

	err := gormDeviceRepository.CreateDevice(gormDeviceOne)

	assert.NoError(t, err)

	gormDeviceTwo := NewGormDevice()

	gormDeviceTwo.SetUUID("2")
	gormDeviceTwo.SetUserUUID("2")

	err = gormDeviceRepository.CreateDevice(gormDeviceTwo)

	assert.NoError(t, err)

	gormDevices := gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceTwo.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)

	err = gormDeviceRepository.
		DeleteDeviceByUUIDAndUserUUID(gormDeviceOne.GetUUID(), gormDeviceOne.GetUserUUID())

	assert.NoError(t, err)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, len(gormDevices), 0)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceTwo.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)
}

func TestGormDeviceRepository_CreateDevice_And_DeleteAllDevicesByUserUUID(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormDeviceRepository, _ := NewGormDeviceRepository(db)

	gormDeviceOne := NewGormDevice()

	gormDeviceOne.SetUUID("1")
	gormDeviceOne.SetUserUUID("1")

	err := gormDeviceRepository.CreateDevice(gormDeviceOne)

	assert.NoError(t, err)

	gormDeviceOne.SetUUID("11")
	gormDeviceOne.SetUserUUID("1")

	err = gormDeviceRepository.CreateDevice(gormDeviceOne)

	assert.NoError(t, err)

	gormDeviceTwo := NewGormDevice()

	gormDeviceTwo.SetUUID("2")
	gormDeviceTwo.SetUserUUID("2")

	err = gormDeviceRepository.CreateDevice(gormDeviceTwo)

	assert.NoError(t, err)

	gormDevices := gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, len(gormDevices), 2)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceTwo.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)

	err = gormDeviceRepository.DeleteAllDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.NoError(t, err)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, len(gormDevices), 0)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceTwo.GetUserUUID())

	assert.Equal(t, len(gormDevices), 1)
}
