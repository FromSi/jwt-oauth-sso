package repositories

import (
	tokens_mocks "github.com/fromsi/jwt-oauth-sso/mocks/tokens"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func TestGormDevice_GenerateAccessToken(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)
	mockAccessToken := tokens_mocks.NewMockAccessToken(mockController)

	mockAccessTokenBuilder.EXPECT().New().Return(mockAccessTokenBuilder).AnyTimes()
	mockAccessTokenBuilder.EXPECT().SetSubject(gomock.Any()).Return(mockAccessTokenBuilder).AnyTimes()
	mockAccessTokenBuilder.EXPECT().SetDeviceUUID(gomock.Any()).Return(mockAccessTokenBuilder).AnyTimes()
	mockAccessTokenBuilder.EXPECT().SetDeviceUserAgent(gomock.Any()).Return(mockAccessTokenBuilder).AnyTimes()
	mockAccessTokenBuilder.EXPECT().SetIssuedAt(gomock.Any()).Return(mockAccessTokenBuilder).AnyTimes()
	mockAccessTokenBuilder.EXPECT().SetExpirationTime(gomock.Any()).Return(mockAccessTokenBuilder).AnyTimes()
	mockAccessTokenBuilder.EXPECT().Build().Return(mockAccessToken, nil).AnyTimes()

	gormDevice := GormDevice{
		accessTokenBuilder: mockAccessTokenBuilder,
	}

	accessToken, err := gormDevice.GenerateAccessToken()

	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)
}

func TestGormDevice_TableName(t *testing.T) {
	gormDevice := GormDevice{}

	assert.Equal(t, gormDevice.TableName(), "devices")
}

func TestGormDevice_GetUUID(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.UUID = "1"

	assert.Equal(t, "1", gormDevice.GetUUID())

	gormDevice.UUID = "2"

	assert.Equal(t, "2", gormDevice.GetUUID())
}

func TestGormDevice_GetUserUUID(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.UserUUID = "1"

	assert.Equal(t, "1", gormDevice.GetUserUUID())

	gormDevice.UserUUID = "2"

	assert.Equal(t, "2", gormDevice.GetUserUUID())
}

func TestGormDevice_GetUserAgent(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.UserAgent = "1"

	assert.Equal(t, "1", gormDevice.GetUserAgent())

	gormDevice.UserAgent = "2"

	assert.Equal(t, "2", gormDevice.GetUserAgent())
}

func TestGormDevice_GetIp(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.Ip = "1"

	assert.Equal(t, "1", gormDevice.GetIp())

	gormDevice.Ip = "2"

	assert.Equal(t, "2", gormDevice.GetIp())
}

func TestGormDevice_GetRefreshToken(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.RefreshToken = "1"

	assert.Equal(t, "1", gormDevice.GetRefreshToken())

	gormDevice.RefreshToken = "2"

	assert.Equal(t, "2", gormDevice.GetRefreshToken())
}

func TestGormDevice_GetIssuedAt(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.IssuedAt = 1

	assert.Equal(t, 1, gormDevice.GetIssuedAt())

	gormDevice.IssuedAt = 2

	assert.Equal(t, 2, gormDevice.GetIssuedAt())
}

func TestGormDevice_GetExpiresAt(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.ExpiresAt = 1

	assert.Equal(t, 1, gormDevice.GetExpiresAt())

	gormDevice.ExpiresAt = 2

	assert.Equal(t, 2, gormDevice.GetExpiresAt())
}

func TestGormDevice_GetCreatedAt(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.CreatedAt = 1

	assert.Equal(t, 1, gormDevice.GetCreatedAt())

	gormDevice.CreatedAt = 2

	assert.Equal(t, 2, gormDevice.GetCreatedAt())
}

func TestGormDevice_GetUpdatedAt(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.UpdatedAt = 1

	assert.Equal(t, 1, gormDevice.GetUpdatedAt())

	gormDevice.UpdatedAt = 2

	assert.Equal(t, 2, gormDevice.GetUpdatedAt())
}

func TestGormDevice_SetUUID(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.SetUUID("1")

	assert.Equal(t, "1", gormDevice.UUID)

	gormDevice.SetUUID("2")

	assert.Equal(t, "2", gormDevice.UUID)
}

func TestGormDevice_SetUserUUID(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.SetUserUUID("1")

	assert.Equal(t, "1", gormDevice.UserUUID)

	gormDevice.SetUserUUID("2")

	assert.Equal(t, "2", gormDevice.UserUUID)
}

func TestGormDevice_SetUserAgent(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.SetUserAgent("1")

	assert.Equal(t, "1", gormDevice.UserAgent)

	gormDevice.SetUserAgent("2")

	assert.Equal(t, "2", gormDevice.UserAgent)
}

func TestGormDevice_SetIp(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.SetIp("1")

	assert.Equal(t, "1", gormDevice.Ip)

	gormDevice.SetIp("2")

	assert.Equal(t, "2", gormDevice.Ip)
}

func TestGormDevice_SetRefreshToken(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.SetRefreshToken("1")

	assert.Equal(t, "1", gormDevice.RefreshToken)

	gormDevice.SetRefreshToken("2")

	assert.Equal(t, "2", gormDevice.RefreshToken)
}

func TestGormDevice_SetIssuedAt(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.SetIssuedAt(1)

	assert.Equal(t, 1, gormDevice.IssuedAt)

	gormDevice.SetIssuedAt(2)

	assert.Equal(t, 2, gormDevice.IssuedAt)
}

func TestGormDevice_SetExpiresAt(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.SetExpiresAt(1)

	assert.Equal(t, 1, gormDevice.ExpiresAt)

	gormDevice.SetExpiresAt(2)

	assert.Equal(t, 2, gormDevice.ExpiresAt)
}

func TestGormDevice_SetCreatedAt(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.SetCreatedAt(1)

	assert.Equal(t, 1, gormDevice.CreatedAt)

	gormDevice.SetCreatedAt(2)

	assert.Equal(t, 2, gormDevice.CreatedAt)
}

func TestGormDevice_SetUpdatedAt(t *testing.T) {
	gormDevice := GormDevice{}

	gormDevice.SetUpdatedAt(1)

	assert.Equal(t, 1, gormDevice.UpdatedAt)

	gormDevice.SetUpdatedAt(2)

	assert.Equal(t, 2, gormDevice.UpdatedAt)
}

func Test_NewGormDeviceRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?mode=ro"), &gorm.Config{})

	assert.NoError(t, err)
	assert.NotEmpty(t, db)

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	deviceBuilder := NewBaseDeviceBuilder(mockAccessTokenBuilder)

	gormDeviceRepository, err := NewGormDeviceRepository(db, deviceBuilder)

	assert.Error(t, err)
	assert.Empty(t, gormDeviceRepository)

	db, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	assert.NoError(t, err)
	assert.NotEmpty(t, db)

	stmt := &gorm.Statement{DB: db}

	var tableCount int

	db.
		Raw(
			"SELECT count(*) FROM sqlite_master WHERE type = 'table' AND name = ?",
			stmt.Table,
		).
		Scan(&tableCount)

	assert.Equal(t, 0, tableCount)

	gormDeviceRepository, err = NewGormDeviceRepository(db, deviceBuilder)

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceRepository)

	err = stmt.Parse(&GormDevice{})

	assert.NoError(t, err)

	db.
		Raw(
			"SELECT count(*) FROM sqlite_master WHERE type = 'table' AND name = ?",
			(&GormDevice{}).TableName(),
		).
		Scan(&tableCount)

	assert.Equal(t, 1, tableCount)
}

func TestGormDeviceRepository_CreateDevice_And_GetDeviceByUserUUIDAndIpAndUserAgent(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	deviceBuilder := NewBaseDeviceBuilder(mockAccessTokenBuilder)

	gormDeviceRepository, err := NewGormDeviceRepository(db, deviceBuilder)

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceRepository)

	gormDevice, err := deviceBuilder.
		New().
		SetUUID("1").
		SetUserUUID("1").
		SetUserAgent("1").
		SetIp("1").
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDevice)

	err = gormDeviceRepository.CreateDevice(gormDevice)

	assert.NoError(t, err)

	err = gormDeviceRepository.CreateDevice(&GormDevice{})

	assert.Error(t, err)

	gormDeviceResult := gormDeviceRepository.
		GetDeviceByUserUUIDAndIpAndUserAgent("0", "0", "0")

	assert.Empty(t, gormDeviceResult)

	gormDeviceResult = gormDeviceRepository.GetDeviceByUserUUIDAndIpAndUserAgent(
		gormDevice.GetUserUUID(),
		gormDevice.GetIp(),
		gormDevice.GetUserAgent(),
	)

	assert.NotEmpty(t, gormDeviceResult)

	assert.Equal(t, gormDevice.GetUserUUID(), gormDeviceResult.GetUserUUID())
	assert.Equal(t, gormDevice.GetUserAgent(), gormDeviceResult.GetUserAgent())
	assert.Equal(t, gormDevice.GetIp(), gormDeviceResult.GetIp())
}

func TestGormDeviceRepository_CreateDevice_And_GetDeviceByRefreshToken(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	deviceBuilder := NewBaseDeviceBuilder(mockAccessTokenBuilder)

	gormDeviceRepository, err := NewGormDeviceRepository(db, deviceBuilder)

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceRepository)

	gormDevice, err := deviceBuilder.
		New().
		SetUUID("1").
		SetRefreshToken("1").
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDevice)

	err = gormDeviceRepository.CreateDevice(gormDevice)

	assert.NoError(t, err)

	gormDeviceResult := gormDeviceRepository.GetDeviceByRefreshToken("0")

	assert.Empty(t, gormDeviceResult)

	gormDeviceResult = gormDeviceRepository.GetDeviceByRefreshToken(
		gormDevice.GetRefreshToken(),
	)

	assert.NotEmpty(t, gormDeviceResult)

	assert.Equal(t, gormDevice.GetRefreshToken(), gormDeviceResult.GetRefreshToken())
}

func TestGormDeviceRepository_CreateDevice_And_GetDevicesByUserUUID_And_UpdateDevice(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	deviceBuilder := NewBaseDeviceBuilder(mockAccessTokenBuilder)

	gormDeviceRepository, err := NewGormDeviceRepository(db, deviceBuilder)

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceRepository)

	gormDeviceOne, err := deviceBuilder.
		New().
		SetUUID("1").
		SetUserUUID("1").
		SetUpdatedAt(1).
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceOne)

	err = gormDeviceRepository.CreateDevice(gormDeviceOne)

	assert.NoError(t, err)

	gormDeviceTwo, err := deviceBuilder.
		New().
		SetUUID("2").
		SetUserUUID("2").
		SetUpdatedAt(2).
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceTwo)

	err = gormDeviceRepository.CreateDevice(gormDeviceTwo)

	assert.NoError(t, err)

	gormDevices := gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, 1, len(gormDevices))
	assert.Equal(t, gormDeviceOne.GetUserUUID(), gormDevices[0].GetUserUUID())
	assert.Equal(t, gormDeviceOne.GetUpdatedAt(), gormDevices[0].GetUpdatedAt())

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceTwo.GetUserUUID())

	assert.Equal(t, 1, len(gormDevices))
	assert.Equal(t, gormDeviceTwo.GetUserUUID(), gormDevices[0].GetUserUUID())
	assert.Equal(t, gormDeviceTwo.GetUpdatedAt(), gormDevices[0].GetUpdatedAt())

	gormDeviceOne.SetUserUUID("3")
	gormDeviceOne.SetUpdatedAt(3)

	err = gormDeviceRepository.UpdateDevice(gormDeviceOne)

	assert.NoError(t, err)

	err = gormDeviceRepository.UpdateDevice(&GormDevice{})

	assert.Error(t, err)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, 1, len(gormDevices))
	assert.Equal(t, gormDeviceOne.GetUserUUID(), gormDevices[0].GetUserUUID())
	assert.Equal(t, gormDeviceOne.GetUpdatedAt(), gormDevices[0].GetUpdatedAt())

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceTwo.GetUserUUID())

	assert.Equal(t, 1, len(gormDevices))
	assert.Equal(t, gormDeviceTwo.GetUserUUID(), gormDevices[0].GetUserUUID())
	assert.Equal(t, gormDeviceTwo.GetUpdatedAt(), gormDevices[0].GetUpdatedAt())
}

func TestGormDeviceRepository_CreateDevice_And_GetDevicesByUserUUID_And_DeleteDeviceByUUID(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	deviceBuilder := NewBaseDeviceBuilder(mockAccessTokenBuilder)

	gormDeviceRepository, err := NewGormDeviceRepository(db, deviceBuilder)

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceRepository)

	gormDeviceOne, err := deviceBuilder.
		New().
		SetUUID("1").
		SetUserUUID("1").
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceOne)

	err = gormDeviceRepository.CreateDevice(gormDeviceOne)

	assert.NoError(t, err)

	gormDeviceTwo, err := deviceBuilder.
		New().
		SetUUID("2").
		SetUserUUID("2").
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceTwo)

	err = gormDeviceRepository.CreateDevice(gormDeviceTwo)

	assert.NoError(t, err)

	gormDevices := gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, 1, len(gormDevices))

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceTwo.GetUserUUID())

	assert.Equal(t, 1, len(gormDevices))

	err = gormDeviceRepository.DeleteDeviceByUUID(gormDeviceOne.GetUUID())

	assert.NoError(t, err)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, 0, len(gormDevices))

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceTwo.GetUserUUID())

	assert.Equal(t, 1, len(gormDevices))
}

func TestGormDeviceRepository_CreateDevice_And_DeleteDeviceByUUIDAndUserUUID(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	deviceBuilder := NewBaseDeviceBuilder(mockAccessTokenBuilder)

	gormDeviceRepository, err := NewGormDeviceRepository(db, deviceBuilder)

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceRepository)

	gormDeviceOne, err := deviceBuilder.
		New().
		SetUUID("1").
		SetUserUUID("1").
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceOne)

	err = gormDeviceRepository.CreateDevice(gormDeviceOne)

	assert.NoError(t, err)

	gormDeviceTwo, err := deviceBuilder.
		New().
		SetUUID("2").
		SetUserUUID("2").
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceTwo)

	err = gormDeviceRepository.CreateDevice(gormDeviceTwo)

	assert.NoError(t, err)

	gormDevices := gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, 1, len(gormDevices))

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceTwo.GetUserUUID())

	assert.Equal(t, 1, len(gormDevices))

	err = gormDeviceRepository.
		DeleteDeviceByUUIDAndUserUUID(gormDeviceOne.GetUUID(), gormDeviceOne.GetUserUUID())

	assert.NoError(t, err)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, 0, len(gormDevices))

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceTwo.GetUserUUID())

	assert.Equal(t, 1, len(gormDevices))
}

func TestGormDeviceRepository_CreateDevice_And_DeleteAllDevicesByUserUUID(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	deviceBuilder := NewBaseDeviceBuilder(mockAccessTokenBuilder)

	gormDeviceRepository, err := NewGormDeviceRepository(db, deviceBuilder)

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceRepository)

	gormDeviceOne, err := deviceBuilder.
		New().
		SetUUID("1").
		SetUserUUID("1").
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceOne)

	err = gormDeviceRepository.CreateDevice(gormDeviceOne)

	assert.NoError(t, err)

	gormDeviceTwo, err := deviceBuilder.
		New().
		SetUUID("2").
		SetUserUUID("1").
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceTwo)

	err = gormDeviceRepository.CreateDevice(gormDeviceTwo)

	assert.NoError(t, err)

	gormDeviceThree, err := deviceBuilder.
		New().
		SetUUID("3").
		SetUserUUID("3").
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormDeviceThree)

	err = gormDeviceRepository.CreateDevice(gormDeviceThree)

	assert.NoError(t, err)

	gormDevices := gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, 2, len(gormDevices))

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceThree.GetUserUUID())

	assert.Equal(t, 1, len(gormDevices))

	err = gormDeviceRepository.DeleteAllDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.NoError(t, err)

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceOne.GetUserUUID())

	assert.Equal(t, 0, len(gormDevices))

	gormDevices = gormDeviceRepository.GetDevicesByUserUUID(gormDeviceThree.GetUserUUID())

	assert.Equal(t, 1, len(gormDevices))
}
