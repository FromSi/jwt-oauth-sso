package repositories

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

func TestGormResetToken_TableName(t *testing.T) {
	gormUser := GormResetToken{}

	assert.Equal(t, gormUser.TableName(), "reset_tokens")
}

func TestGormResetToken_GetToken(t *testing.T) {
	gormResetToken := GormResetToken{}

	gormResetToken.Token = "1"

	assert.Equal(t, "1", gormResetToken.GetToken())

	gormResetToken.Token = "2"

	assert.Equal(t, "2", gormResetToken.GetToken())
}

func TestGormResetToken_GetUserUUID(t *testing.T) {
	gormResetToken := GormResetToken{}

	gormResetToken.UserUUID = "1"

	assert.Equal(t, "1", gormResetToken.GetUserUUID())

	gormResetToken.UserUUID = "2"

	assert.Equal(t, "2", gormResetToken.GetUserUUID())
}

func TestGormResetToken_GetExpiresAt(t *testing.T) {
	gormResetToken := GormResetToken{}

	gormResetToken.ExpiresAt = 1

	assert.Equal(t, 1, gormResetToken.GetExpiresAt())

	gormResetToken.ExpiresAt = 2

	assert.Equal(t, 2, gormResetToken.GetExpiresAt())
}

func TestGormResetToken_GetCreatedAt(t *testing.T) {
	gormResetToken := GormResetToken{}

	gormResetToken.CreatedAt = 1

	assert.Equal(t, 1, gormResetToken.GetCreatedAt())

	gormResetToken.CreatedAt = 2

	assert.Equal(t, 2, gormResetToken.GetCreatedAt())
}

func TestGormResetToken_SetToken(t *testing.T) {
	gormResetToken := GormResetToken{}

	gormResetToken.SetToken("1")

	assert.Equal(t, "1", gormResetToken.Token)

	gormResetToken.SetToken("2")

	assert.Equal(t, "2", gormResetToken.Token)
}

func TestGormResetToken_SetUserUUID(t *testing.T) {
	gormResetToken := GormResetToken{}

	gormResetToken.SetUserUUID("1")

	assert.Equal(t, "1", gormResetToken.UserUUID)

	gormResetToken.SetUserUUID("2")

	assert.Equal(t, "2", gormResetToken.UserUUID)
}

func TestGormResetToken_SetExpiresAt(t *testing.T) {
	gormResetToken := GormResetToken{}

	gormResetToken.SetExpiresAt(1)

	assert.Equal(t, 1, gormResetToken.ExpiresAt)

	gormResetToken.SetExpiresAt(2)

	assert.Equal(t, 2, gormResetToken.ExpiresAt)
}

func TestGormResetToken_SetCreatedAt(t *testing.T) {
	gormResetToken := GormResetToken{}

	gormResetToken.SetCreatedAt(1)

	assert.Equal(t, 1, gormResetToken.CreatedAt)

	gormResetToken.SetCreatedAt(2)

	assert.Equal(t, 2, gormResetToken.CreatedAt)
}

func Test_NewGormResetTokenRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?mode=ro"), &gorm.Config{})

	assert.NoError(t, err)
	assert.NotEmpty(t, db)

	resetTokenBuilder := NewBaseResetTokenBuilder()

	gormResetTokenRepository, err := NewGormResetTokenRepository(db, resetTokenBuilder)

	assert.Error(t, err)
	assert.Empty(t, gormResetTokenRepository)

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

	gormResetTokenRepository, err = NewGormResetTokenRepository(db, resetTokenBuilder)

	assert.NoError(t, err)
	assert.NotEmpty(t, gormResetTokenRepository)

	err = stmt.Parse(&GormResetToken{})

	assert.NoError(t, err)

	db.
		Raw(
			"SELECT count(*) FROM sqlite_master WHERE type = 'table' AND name = ?",
			(&GormResetToken{}).TableName(),
		).
		Scan(&tableCount)

	assert.Equal(t, 1, tableCount)
}

func TestGormResetTokenRepository_CreateToken_And_GetActiveResetTokenByToken(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	resetTokenBuilder := NewBaseResetTokenBuilder()

	gormResetTokenRepository, err := NewGormResetTokenRepository(db, resetTokenBuilder)

	assert.NoError(t, err)
	assert.NotEmpty(t, gormResetTokenRepository)

	gormResetTokenOne, err := resetTokenBuilder.
		New().
		SetToken("1").
		SetExpiresAt(int(time.Now().AddDate(0, 0, 1).Unix())).
		Build()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormResetTokenOne)

	err = gormResetTokenRepository.CreateResetToken(gormResetTokenOne)

	assert.NoError(t, err)

	err = gormResetTokenRepository.CreateResetToken(&GormResetToken{})

	assert.Error(t, err)

	gormResetTokenTwo, err := resetTokenBuilder.
		New().
		SetToken("2").
		SetExpiresAt(0).
		Build()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormResetTokenTwo)

	err = gormResetTokenRepository.CreateResetToken(gormResetTokenTwo)

	assert.NoError(t, err)

	resetTokenByToken := gormResetTokenRepository.
		GetActiveResetTokenByToken(gormResetTokenOne.GetToken())

	assert.NotEmpty(t, resetTokenByToken)

	assert.Equal(t, gormResetTokenOne.GetToken(), resetTokenByToken.GetToken())

	resetTokenByToken = gormResetTokenRepository.GetActiveResetTokenByToken("2")

	assert.Empty(t, resetTokenByToken)
}

func TestGormResetTokenRepository_CreateToken_And_DeleteResetToken(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	resetTokenBuilder := NewBaseResetTokenBuilder()

	gormResetTokenRepository, err := NewGormResetTokenRepository(db, resetTokenBuilder)

	assert.NoError(t, err)
	assert.NotEmpty(t, gormResetTokenRepository)

	gormResetTokenOne, err := resetTokenBuilder.
		New().
		SetToken("1").
		SetExpiresAt(int(time.Now().AddDate(0, 0, 1).Unix())).
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormResetTokenOne)

	err = gormResetTokenRepository.CreateResetToken(gormResetTokenOne)

	assert.NoError(t, err)

	gormResetTokenTwo, err := resetTokenBuilder.
		New().
		SetToken("2").
		SetExpiresAt(int(time.Now().AddDate(0, 0, 1).Unix())).
		BuildToGorm()

	assert.NoError(t, err)
	assert.NotEmpty(t, gormResetTokenTwo)

	err = gormResetTokenRepository.CreateResetToken(gormResetTokenTwo)

	assert.NoError(t, err)

	result := gormResetTokenRepository.GetActiveResetTokenByToken(gormResetTokenOne.GetToken())

	assert.NotEmpty(t, result)

	result = gormResetTokenRepository.GetActiveResetTokenByToken(gormResetTokenTwo.GetToken())

	assert.NotEmpty(t, result)

	err = gormResetTokenRepository.DeleteResetToken(gormResetTokenOne.GetToken())

	assert.NoError(t, err)

	result = gormResetTokenRepository.GetActiveResetTokenByToken(gormResetTokenOne.GetToken())

	assert.Empty(t, result)

	result = gormResetTokenRepository.GetActiveResetTokenByToken(gormResetTokenTwo.GetToken())

	assert.NotEmpty(t, result)
}
