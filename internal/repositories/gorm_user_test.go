package repositories

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func Test_NewGormUser(t *testing.T) {
	gormUser := NewGormUser()

	assert.NotNil(t, gormUser)

	assert.Equal(t, gormUser.UUID, GormUserUUIDDefault)
	assert.Equal(t, gormUser.Email, GormUserEmailDefault)
	assert.Equal(t, gormUser.Password, GormUserPasswordDefault)
	assert.Equal(t, gormUser.CreatedAt, GormUserCreatedAtDefault)
	assert.Equal(t, gormUser.UpdatedAt, GormUserUpdatedAtDefault)
}

func Test_NewGormUserByUser(t *testing.T) {
	gormUserTemp := NewGormUser()
	gormUser := NewGormUserByUser(gormUserTemp)

	assert.NotNil(t, gormUser)

	assert.Equal(t, gormUser.UUID, GormUserUUIDDefault)
	assert.Equal(t, gormUser.Email, GormUserEmailDefault)
	assert.Equal(t, gormUser.Password, GormUserPasswordDefault)
	assert.Equal(t, gormUser.CreatedAt, GormUserCreatedAtDefault)
	assert.Equal(t, gormUser.UpdatedAt, GormUserUpdatedAtDefault)
}

func TestGormUser_GetUUID(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.UUID = "1"

	assert.Equal(t, gormUser.GetUUID(), "1")
}

func TestGormUser_GetEmail(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.Email = "1"

	assert.Equal(t, gormUser.GetEmail(), "1")
}

func TestGormUser_GetPassword(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.Password = "1"

	assert.Equal(t, gormUser.GetPassword(), "1")
}

func TestGormUser_GetCreatedAt(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.CreatedAt = 1

	assert.Equal(t, gormUser.GetCreatedAt(), 1)
}

func TestGormUser_GetUpdatedAt(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.UpdatedAt = 1

	assert.Equal(t, gormUser.GetUpdatedAt(), 1)
}

func TestGormUser_SetUUID(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.SetUUID("1")

	assert.Equal(t, gormUser.UUID, "1")
}

func TestGormUser_SetEmail(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.SetEmail("1")

	assert.Equal(t, gormUser.Email, "1")
}

func TestGormUser_SetPassword(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.SetPassword("1")

	assert.Equal(t, gormUser.Password, "1")
}

func TestGormUser_SetCreatedAt(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.SetCreatedAt(1)

	assert.Equal(t, gormUser.CreatedAt, 1)
}

func TestGormUser_SetUpdatedAt(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.SetUpdatedAt(1)

	assert.Equal(t, gormUser.UpdatedAt, 1)
}

func Test_NewGormUserRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	assert.Nil(t, err)
	assert.NotNil(t, db)

	gormUserRepository, err := NewGormUserRepository(db)

	assert.Nil(t, err)
	assert.NotNil(t, gormUserRepository)

	stmt := &gorm.Statement{DB: db}
	err = stmt.Parse(&GormUser{})

	assert.Nil(t, err)

	var count int

	db.Raw("SELECT count(*) FROM sqlite_master WHERE type = 'table' AND name = ?", stmt.Table).Scan(&count)

	assert.Equal(t, count, 1)
}

func TestGormUserRepository_CreateUser_And_HasUserByUUID(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	gormUserRepository, _ := NewGormUserRepository(db)

	gormUser := NewGormUser()

	gormUser.SetUUID("1")

	err := gormUserRepository.CreateUser(gormUser)

	assert.Nil(t, err)

	exists := gormUserRepository.HasUserByUUID(gormUser.GetUUID())

	assert.True(t, exists)

	exists = gormUserRepository.HasUserByUUID("0")

	assert.False(t, exists)
}

func TestGormUserRepository_CreateUser_And_HasUserByEmail(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	gormUserRepository, _ := NewGormUserRepository(db)

	gormUser := NewGormUser()

	gormUser.SetEmail("1")

	err := gormUserRepository.CreateUser(gormUser)

	assert.Nil(t, err)

	exists := gormUserRepository.HasUserByEmail(gormUser.GetEmail())

	assert.True(t, exists)

	exists = gormUserRepository.HasUserByEmail("0")

	assert.False(t, exists)
}

func TestGormUserRepository_CreateUser_And_HasUserByEmailAndPassword(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	gormUserRepository, _ := NewGormUserRepository(db)

	gormUser := NewGormUser()

	gormUser.SetEmail("1")
	gormUser.SetEmail("2")

	err := gormUserRepository.CreateUser(gormUser)

	assert.Nil(t, err)

	exists := gormUserRepository.HasUserByEmailAndPassword(gormUser.GetEmail(), gormUser.GetPassword())

	assert.True(t, exists)

	exists = gormUserRepository.HasUserByEmailAndPassword("0", "0")

	assert.False(t, exists)
}

func TestGormUserRepository_CreateUser_And_HasUserByUUIDAndPassword(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	gormUserRepository, _ := NewGormUserRepository(db)

	gormUser := NewGormUser()

	gormUser.SetUUID("1")
	gormUser.SetPassword("2")

	err := gormUserRepository.CreateUser(gormUser)

	assert.Nil(t, err)

	exists := gormUserRepository.HasUserByUUIDAndPassword(gormUser.GetUUID(), gormUser.GetPassword())

	assert.True(t, exists)

	exists = gormUserRepository.HasUserByUUIDAndPassword("0", "0")

	assert.False(t, exists)
}

func TestGormUserRepository_CreateUser_And_GetUserByEmail(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	gormUserRepository, _ := NewGormUserRepository(db)

	gormUser := NewGormUser()

	gormUser.SetUUID("1")
	gormUser.SetEmail("2")
	gormUser.SetPassword("3")
	gormUser.SetCreatedAt(4)
	gormUser.SetUpdatedAt(5)

	err := gormUserRepository.CreateUser(gormUser)

	assert.Nil(t, err)

	gormUserForRepository := gormUserRepository.GetUserByEmail("0")

	assert.Nil(t, gormUserForRepository)

	gormUserForRepository = gormUserRepository.GetUserByEmail(gormUser.GetEmail())

	assert.NotNil(t, gormUserForRepository)
	assert.Equal(t, gormUserForRepository.GetUUID(), gormUser.GetUUID())
	assert.Equal(t, gormUserForRepository.GetEmail(), gormUser.GetEmail())
	assert.Equal(t, gormUserForRepository.GetPassword(), gormUser.GetPassword())
	assert.Equal(t, gormUserForRepository.GetCreatedAt(), gormUser.GetCreatedAt())
	assert.Equal(t, gormUserForRepository.GetUpdatedAt(), gormUser.GetUpdatedAt())
}

func TestGormUserRepository_CreateUser_And_UpdatePassword(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	gormUserRepository, _ := NewGormUserRepository(db)

	gormUser := NewGormUser()

	gormUser.SetUUID("1")
	gormUser.SetEmail("2")
	gormUser.SetPassword("3")
	gormUser.SetUpdatedAt(4)

	err := gormUserRepository.CreateUser(gormUser)

	assert.Nil(t, err)

	gormUserForRepository := gormUserRepository.GetUserByEmail(gormUser.GetEmail())

	assert.NotNil(t, gormUserForRepository)
	assert.Equal(t, gormUserForRepository.GetPassword(), gormUser.GetPassword())
	assert.Equal(t, gormUserForRepository.GetUpdatedAt(), gormUser.GetUpdatedAt())

	gormUser.SetPassword("5")
	gormUser.SetUpdatedAt(6)

	err = gormUserRepository.UpdatePassword(gormUser.GetUUID(), gormUser.GetPassword(), gormUser.GetUpdatedAt())

	assert.Nil(t, err)

	gormUserForRepository = gormUserRepository.GetUserByEmail(gormUser.GetEmail())

	assert.NotNil(t, gormUserForRepository)
	assert.Equal(t, gormUserForRepository.GetPassword(), gormUser.GetPassword())
	assert.Equal(t, gormUserForRepository.GetUpdatedAt(), gormUser.GetUpdatedAt())
}
