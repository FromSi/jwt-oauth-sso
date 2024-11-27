package repositories

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func TestGormUser_TableName(t *testing.T) {
	gormUser := GormUser{}

	assert.Equal(t, gormUser.TableName(), "users")
}

func TestGormUser_GetUUID(t *testing.T) {
	gormUser := GormUser{}

	gormUser.UUID = "1"

	assert.Equal(t, "1", gormUser.GetUUID())

	gormUser.UUID = "2"

	assert.Equal(t, "2", gormUser.GetUUID())
}

func TestGormUser_GetEmail(t *testing.T) {
	gormUser := GormUser{}

	gormUser.Email = "1"

	assert.Equal(t, "1", gormUser.GetEmail())

	gormUser.Email = "2"

	assert.Equal(t, "2", gormUser.GetEmail())
}

func TestGormUser_GetPassword(t *testing.T) {
	gormUser := GormUser{}

	gormUser.Password = "1"

	assert.Equal(t, "1", gormUser.GetPassword())

	gormUser.Password = "2"

	assert.Equal(t, "2", gormUser.GetPassword())
}

func TestGormUser_GetCreatedAt(t *testing.T) {
	gormUser := GormUser{}

	gormUser.CreatedAt = 1

	assert.Equal(t, 1, gormUser.GetCreatedAt())

	gormUser.CreatedAt = 2

	assert.Equal(t, 2, gormUser.GetCreatedAt())
}

func TestGormUser_GetUpdatedAt(t *testing.T) {
	gormUser := GormUser{}

	gormUser.UpdatedAt = 1

	assert.Equal(t, 1, gormUser.GetUpdatedAt())

	gormUser.UpdatedAt = 2

	assert.Equal(t, 2, gormUser.GetUpdatedAt())
}

func TestGormUser_SetUUID(t *testing.T) {
	gormUser := GormUser{}

	gormUser.SetUUID("1")

	assert.Equal(t, "1", gormUser.UUID)

	gormUser.SetUUID("2")

	assert.Equal(t, "2", gormUser.UUID)
}

func TestGormUser_SetEmail(t *testing.T) {
	gormUser := GormUser{}

	gormUser.SetEmail("1")

	assert.Equal(t, "1", gormUser.Email)

	gormUser.SetEmail("2")

	assert.Equal(t, "2", gormUser.Email)
}

func TestGormUser_SetPassword(t *testing.T) {
	gormUser := GormUser{}

	gormUser.SetPassword("1")

	assert.Equal(t, "1", gormUser.Password)

	gormUser.SetPassword("2")

	assert.Equal(t, "2", gormUser.Password)
}

func TestGormUser_SetCreatedAt(t *testing.T) {
	gormUser := GormUser{}

	gormUser.SetCreatedAt(1)

	assert.Equal(t, 1, gormUser.CreatedAt)

	gormUser.SetCreatedAt(2)

	assert.Equal(t, 2, gormUser.CreatedAt)
}

func TestGormUser_SetUpdatedAt(t *testing.T) {
	gormUser := GormUser{}

	gormUser.SetUpdatedAt(1)

	assert.Equal(t, 1, gormUser.UpdatedAt)

	gormUser.SetUpdatedAt(2)

	assert.Equal(t, 2, gormUser.UpdatedAt)
}

func Test_NewGormUserRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?mode=ro"), &gorm.Config{})

	assert.NoError(t, err)
	assert.NotEmpty(t, db)

	userBuilder := NewBaseUserBuilder()

	gormUserRepository, err := NewGormUserRepository(db, userBuilder)

	assert.Error(t, err)
	assert.Empty(t, gormUserRepository)

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

	gormUserRepository, err = NewGormUserRepository(db, userBuilder)

	assert.NoError(t, err)
	assert.NotEmpty(t, gormUserRepository)

	err = stmt.Parse(&GormUser{})

	assert.NoError(t, err)

	db.
		Raw(
			"SELECT count(*) FROM sqlite_master WHERE type = 'table' AND name = ?",
			(&GormUser{}).TableName(),
		).
		Scan(&tableCount)

	assert.Equal(t, 1, tableCount)
}

func TestGormUserRepository_CreateUser_And_GetUserByEmail(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	userBuilder := NewBaseUserBuilder()

	gormUserRepository, err := NewGormUserRepository(db, userBuilder)

	assert.NoError(t, err)
	assert.NotEmpty(t, gormUserRepository)

	gormUser, err := userBuilder.
		New().
		SetUUID("1").
		SetEmail("1").
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormUser)

	err = gormUserRepository.CreateUser(gormUser)

	assert.NoError(t, err)

	err = gormUserRepository.CreateUser(&GormUser{})

	assert.Error(t, err)

	gormUserForRepository := gormUserRepository.GetUserByEmail(gormUser.GetEmail())

	assert.NotEmpty(t, gormUserForRepository)
	assert.Equal(t, gormUser.GetEmail(), gormUserForRepository.GetEmail())

	gormUserForRepository = gormUserRepository.GetUserByEmail("0")

	assert.Empty(t, gormUserForRepository)
}

func TestGormUserRepository_CreateUser_And_GetUserByUUID(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	userBuilder := NewBaseUserBuilder()

	gormUserRepository, err := NewGormUserRepository(db, userBuilder)

	assert.NoError(t, err)
	assert.NotEmpty(t, gormUserRepository)

	gormUser, err := userBuilder.
		New().
		SetUUID("1").
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormUser)

	err = gormUserRepository.CreateUser(gormUser)

	assert.NoError(t, err)

	gormUserForRepository := gormUserRepository.GetUserByUUID(gormUser.GetUUID())

	assert.NotEmpty(t, gormUserForRepository)
	assert.Equal(t, gormUser.GetUUID(), gormUserForRepository.GetUUID())

	gormUserForRepository = gormUserRepository.GetUserByUUID("0")

	assert.Empty(t, gormUserForRepository)
}

func TestGormUserRepository_CreateUser_And_UpdatePasswordByUUIDAndPasswordAndUpdatedAt(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	userBuilder := NewBaseUserBuilder()

	gormUserRepository, _ := NewGormUserRepository(db, userBuilder)

	gormUserOne, err := userBuilder.
		New().
		SetUUID("1").
		SetPassword("1").
		SetUpdatedAt(1).
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormUserOne)

	err = gormUserRepository.CreateUser(gormUserOne)

	assert.NoError(t, err)

	gormUserTwo, err := userBuilder.
		New().
		SetUUID("2").
		SetPassword("2").
		SetUpdatedAt(2).
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormUserTwo)

	err = gormUserRepository.CreateUser(gormUserTwo)

	assert.NoError(t, err)

	gormUserForRepository := gormUserRepository.GetUserByUUID(gormUserOne.GetUUID())

	assert.NotEmpty(t, gormUserForRepository)

	assert.Equal(t, gormUserOne.GetPassword(), gormUserForRepository.GetPassword())
	assert.Equal(t, gormUserOne.GetUpdatedAt(), gormUserForRepository.GetUpdatedAt())

	gormUserForRepository = gormUserRepository.GetUserByUUID(gormUserTwo.GetUUID())

	assert.NotEmpty(t, gormUserForRepository)

	assert.Equal(t, gormUserTwo.GetPassword(), gormUserForRepository.GetPassword())
	assert.Equal(t, gormUserTwo.GetUpdatedAt(), gormUserForRepository.GetUpdatedAt())

	gormUserOne.SetPassword("11")
	gormUserOne.SetUpdatedAt(11)

	err = gormUserRepository.UpdatePasswordByUUIDAndPasswordAndUpdatedAt(
		gormUserOne.GetUUID(),
		gormUserOne.GetPassword(),
		gormUserOne.GetUpdatedAt(),
	)

	assert.NoError(t, err)

	gormUserForRepository = gormUserRepository.GetUserByUUID(gormUserOne.GetUUID())

	assert.NotEmpty(t, gormUserForRepository)

	assert.Equal(t, gormUserOne.GetPassword(), gormUserForRepository.GetPassword())
	assert.Equal(t, gormUserOne.GetUpdatedAt(), gormUserForRepository.GetUpdatedAt())

	gormUserForRepository = gormUserRepository.GetUserByUUID(gormUserTwo.GetUUID())

	assert.NotEmpty(t, gormUserForRepository)

	assert.Equal(t, gormUserTwo.GetPassword(), gormUserForRepository.GetPassword())
	assert.Equal(t, gormUserTwo.GetUpdatedAt(), gormUserForRepository.GetUpdatedAt())
}
