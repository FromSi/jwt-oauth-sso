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

	gormUser.UUID = "2"

	assert.Equal(t, gormUser.GetUUID(), "2")
}

func TestGormUser_GetEmail(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.Email = "1"

	assert.Equal(t, gormUser.GetEmail(), "1")

	gormUser.Email = "2"

	assert.Equal(t, gormUser.GetEmail(), "2")
}

func TestGormUser_GetPassword(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.Password = "1"

	assert.Equal(t, gormUser.GetPassword(), "1")

	gormUser.Password = "2"

	assert.Equal(t, gormUser.GetPassword(), "2")
}

func TestGormUser_GetCreatedAt(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.CreatedAt = 1

	assert.Equal(t, gormUser.GetCreatedAt(), 1)

	gormUser.CreatedAt = 2

	assert.Equal(t, gormUser.GetCreatedAt(), 2)
}

func TestGormUser_GetUpdatedAt(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.UpdatedAt = 1

	assert.Equal(t, gormUser.GetUpdatedAt(), 1)

	gormUser.UpdatedAt = 2

	assert.Equal(t, gormUser.GetUpdatedAt(), 2)
}

func TestGormUser_SetUUID(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.SetUUID("1")

	assert.Equal(t, gormUser.UUID, "1")

	gormUser.SetUUID("2")

	assert.Equal(t, gormUser.UUID, "2")
}

func TestGormUser_SetEmail(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.SetEmail("1")

	assert.Equal(t, gormUser.Email, "1")

	gormUser.SetEmail("2")

	assert.Equal(t, gormUser.Email, "2")
}

func TestGormUser_SetPassword(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.SetPassword("1")

	assert.Equal(t, gormUser.Password, "1")

	gormUser.SetPassword("2")

	assert.Equal(t, gormUser.Password, "2")
}

func TestGormUser_SetCreatedAt(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.SetCreatedAt(1)

	assert.Equal(t, gormUser.CreatedAt, 1)

	gormUser.SetCreatedAt(2)

	assert.Equal(t, gormUser.CreatedAt, 2)
}

func TestGormUser_SetUpdatedAt(t *testing.T) {
	gormUser := NewGormUser()

	gormUser.SetUpdatedAt(1)

	assert.Equal(t, gormUser.UpdatedAt, 1)

	gormUser.SetUpdatedAt(2)

	assert.Equal(t, gormUser.UpdatedAt, 2)
}

func Test_NewGormUserRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?mode=ro"), &gorm.Config{})

	assert.Nil(t, err)
	assert.NotNil(t, db)

	gormUserRepository, err := NewGormUserRepository(db)

	assert.NotNil(t, err)
	assert.Nil(t, gormUserRepository)

	db, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	assert.Nil(t, err)
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

	gormUserRepository, err = NewGormUserRepository(db)

	assert.Nil(t, err)
	assert.NotNil(t, gormUserRepository)

	err = stmt.Parse(&GormUser{})

	assert.Nil(t, err)

	db.
		Raw(
			"SELECT count(*) FROM sqlite_master WHERE type = 'table' AND name = ?",
			(&GormUser{}).TableName(),
		).
		Scan(&tableCount)

	assert.Equal(t, tableCount, 1)
}

func TestGormUserRepository_CreateUser_And_GetUserByEmail(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormUserRepository, _ := NewGormUserRepository(db)

	gormUser := NewGormUser()

	gormUser.SetEmail("1")

	err := gormUserRepository.CreateUser(gormUser)

	assert.Nil(t, err)

	gormUserForRepository := gormUserRepository.GetUserByEmail(gormUser.GetEmail())

	assert.NotNil(t, gormUserForRepository)
	assert.Equal(t, gormUserForRepository.GetEmail(), gormUser.GetEmail())

	gormUserForRepository = gormUserRepository.GetUserByEmail("0")

	assert.Nil(t, gormUserForRepository)
}

func TestGormUserRepository_CreateUser_And_GetUserByUUID(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormUserRepository, _ := NewGormUserRepository(db)

	gormUser := NewGormUser()

	gormUser.SetUUID("1")

	err := gormUserRepository.CreateUser(gormUser)

	assert.Nil(t, err)

	gormUserForRepository := gormUserRepository.GetUserByUUID(gormUser.GetUUID())

	assert.NotNil(t, gormUserForRepository)
	assert.Equal(t, gormUserForRepository.GetUUID(), gormUser.GetUUID())

	gormUserForRepository = gormUserRepository.GetUserByUUID("0")

	assert.Nil(t, gormUserForRepository)
}

func TestGormUserRepository_CreateUser_And_UpdatePasswordByUUIDAndPasswordAndUpdatedAt(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormUserRepository, _ := NewGormUserRepository(db)

	gormUserOne := NewGormUser()

	gormUserOne.SetUUID("1")
	gormUserOne.SetPassword("1")
	gormUserOne.SetUpdatedAt(1)

	err := gormUserRepository.CreateUser(gormUserOne)

	assert.Nil(t, err)

	gormUserTwo := NewGormUser()

	gormUserTwo.SetUUID("2")
	gormUserTwo.SetPassword("2")
	gormUserTwo.SetUpdatedAt(2)

	err = gormUserRepository.CreateUser(gormUserTwo)

	assert.Nil(t, err)

	gormUserForRepository := gormUserRepository.GetUserByUUID(gormUserOne.GetUUID())

	assert.NotNil(t, gormUserForRepository)

	assert.Equal(t, gormUserForRepository.GetPassword(), gormUserOne.GetPassword())
	assert.Equal(t, gormUserForRepository.GetUpdatedAt(), gormUserOne.GetUpdatedAt())

	gormUserForRepository = gormUserRepository.GetUserByUUID(gormUserTwo.GetUUID())

	assert.NotNil(t, gormUserForRepository)

	assert.Equal(t, gormUserForRepository.GetPassword(), gormUserTwo.GetPassword())
	assert.Equal(t, gormUserForRepository.GetUpdatedAt(), gormUserTwo.GetUpdatedAt())

	gormUserOne.SetPassword("11")
	gormUserOne.SetUpdatedAt(11)

	err = gormUserRepository.UpdatePasswordByUUIDAndPasswordAndUpdatedAt(
		gormUserOne.GetUUID(),
		gormUserOne.GetPassword(),
		gormUserOne.GetUpdatedAt(),
	)

	assert.Nil(t, err)

	gormUserForRepository = gormUserRepository.GetUserByUUID(gormUserOne.GetUUID())

	assert.NotNil(t, gormUserForRepository)

	assert.Equal(t, gormUserForRepository.GetPassword(), gormUserOne.GetPassword())
	assert.Equal(t, gormUserForRepository.GetUpdatedAt(), gormUserOne.GetUpdatedAt())

	gormUserForRepository = gormUserRepository.GetUserByUUID(gormUserTwo.GetUUID())

	assert.NotNil(t, gormUserForRepository)

	assert.Equal(t, gormUserForRepository.GetPassword(), gormUserTwo.GetPassword())
	assert.Equal(t, gormUserForRepository.GetUpdatedAt(), gormUserTwo.GetUpdatedAt())
}
