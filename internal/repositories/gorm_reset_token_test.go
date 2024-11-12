package repositories

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

func Test_NewGormResetToken(t *testing.T) {
	gormResetToken := NewGormResetToken()

	assert.NotNil(t, gormResetToken)

	assert.Equal(t, gormResetToken.Token, GormResetTokenTokenDefault)
	assert.Equal(t, gormResetToken.UserUUID, GormResetTokenUserUUIDDefault)
	assert.Equal(t, gormResetToken.ExpiresAt, GormResetTokenExpiresAtDefault)
	assert.Equal(t, gormResetToken.CreatedAt, GormResetTokenCreatedAtDefault)
}

func Test_NewGormResetTokenByResetToken(t *testing.T) {
	gormResetTokenTemp := NewGormResetToken()
	gormResetToken := NewGormResetTokenByResetToken(gormResetTokenTemp)

	assert.NotNil(t, gormResetToken)

	assert.Equal(t, gormResetToken.Token, GormResetTokenTokenDefault)
	assert.Equal(t, gormResetToken.UserUUID, GormResetTokenUserUUIDDefault)
	assert.Equal(t, gormResetToken.ExpiresAt, GormResetTokenExpiresAtDefault)
	assert.Equal(t, gormResetToken.CreatedAt, GormResetTokenCreatedAtDefault)

}

func TestGormResetToken_GetToken(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.Token = "1"

	assert.Equal(t, gormResetToken.GetToken(), "1")

	gormResetToken.Token = "2"

	assert.Equal(t, gormResetToken.GetToken(), "2")
}

func TestGormResetToken_GetUserUUID(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.UserUUID = "1"

	assert.Equal(t, gormResetToken.GetUserUUID(), "1")

	gormResetToken.UserUUID = "2"

	assert.Equal(t, gormResetToken.GetUserUUID(), "2")
}

func TestGormResetToken_GetExpiresAt(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.ExpiresAt = 1

	assert.Equal(t, gormResetToken.GetExpiresAt(), 1)

	gormResetToken.ExpiresAt = 2

	assert.Equal(t, gormResetToken.GetExpiresAt(), 2)
}

func TestGormResetToken_GetCreatedAt(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.CreatedAt = 1

	assert.Equal(t, gormResetToken.GetCreatedAt(), 1)

	gormResetToken.CreatedAt = 2

	assert.Equal(t, gormResetToken.GetCreatedAt(), 2)
}

func TestGormResetToken_SetToken(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.SetToken("1")

	assert.Equal(t, gormResetToken.Token, "1")

	gormResetToken.SetToken("2")

	assert.Equal(t, gormResetToken.Token, "2")
}

func TestGormResetToken_SetUserUUID(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.SetUserUUID("1")

	assert.Equal(t, gormResetToken.UserUUID, "1")

	gormResetToken.SetUserUUID("2")

	assert.Equal(t, gormResetToken.UserUUID, "2")
}

func TestGormResetToken_SetExpiresAt(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.SetExpiresAt(1)

	assert.Equal(t, gormResetToken.ExpiresAt, 1)

	gormResetToken.SetExpiresAt(2)

	assert.Equal(t, gormResetToken.ExpiresAt, 2)
}

func TestGormResetToken_SetCreatedAt(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.SetCreatedAt(1)

	assert.Equal(t, gormResetToken.CreatedAt, 1)

	gormResetToken.SetCreatedAt(2)

	assert.Equal(t, gormResetToken.CreatedAt, 2)
}

func Test_NewGormResetTokenRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?mode=ro"), &gorm.Config{})

	assert.NoError(t, err)
	assert.NotNil(t, db)

	gormResetTokenRepository, err := NewGormResetTokenRepository(db)

	assert.Error(t, err)
	assert.Nil(t, gormResetTokenRepository)

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

	gormResetTokenRepository, err = NewGormResetTokenRepository(db)

	assert.NoError(t, err)
	assert.NotNil(t, gormResetTokenRepository)

	err = stmt.Parse(&GormResetToken{})

	assert.NoError(t, err)

	db.
		Raw(
			"SELECT count(*) FROM sqlite_master WHERE type = 'table' AND name = ?",
			(&GormResetToken{}).TableName(),
		).
		Scan(&tableCount)

	assert.Equal(t, tableCount, 1)
}

func TestGormResetTokenRepository_CreateToken_And_GetActiveResetTokenByToken(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormResetTokenRepository, _ := NewGormResetTokenRepository(db)

	gormResetTokenOne := NewGormResetToken()

	gormResetTokenOne.SetToken("1")
	gormResetTokenOne.SetExpiresAt(int(time.Now().AddDate(0, 0, 1).Unix()))

	err := gormResetTokenRepository.CreateResetToken(gormResetTokenOne)

	assert.NoError(t, err)

	gormResetTokenTwo := NewGormResetToken()

	gormResetTokenTwo.SetToken("2")
	gormResetTokenTwo.SetExpiresAt(0)

	err = gormResetTokenRepository.CreateResetToken(gormResetTokenTwo)

	assert.NoError(t, err)

	resetTokenByToken := gormResetTokenRepository.
		GetActiveResetTokenByToken(gormResetTokenOne.GetToken())

	assert.NotNil(t, resetTokenByToken)

	assert.Equal(t, resetTokenByToken.GetToken(), gormResetTokenOne.GetToken())

	resetTokenByToken = gormResetTokenRepository.GetActiveResetTokenByToken("2")

	assert.Nil(t, resetTokenByToken)
}

func TestGormResetTokenRepository_CreateToken_And_DeleteResetToken(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormResetTokenRepository, _ := NewGormResetTokenRepository(db)

	gormResetTokenOne := NewGormResetToken()

	gormResetTokenOne.SetToken("1")
	gormResetTokenOne.SetExpiresAt(int(time.Now().AddDate(0, 0, 1).Unix()))

	err := gormResetTokenRepository.CreateResetToken(gormResetTokenOne)

	assert.NoError(t, err)

	gormResetTokenTwo := NewGormResetToken()

	gormResetTokenTwo.SetToken("2")
	gormResetTokenTwo.SetExpiresAt(int(time.Now().AddDate(0, 0, 1).Unix()))

	err = gormResetTokenRepository.CreateResetToken(gormResetTokenTwo)

	assert.NoError(t, err)

	result := gormResetTokenRepository.GetActiveResetTokenByToken(gormResetTokenOne.GetToken())

	assert.NotNil(t, result)

	result = gormResetTokenRepository.GetActiveResetTokenByToken(gormResetTokenTwo.GetToken())

	assert.NotNil(t, result)

	err = gormResetTokenRepository.DeleteResetToken(gormResetTokenOne.GetToken())

	assert.NoError(t, err)

	result = gormResetTokenRepository.GetActiveResetTokenByToken(gormResetTokenOne.GetToken())

	assert.Nil(t, result)

	result = gormResetTokenRepository.GetActiveResetTokenByToken(gormResetTokenTwo.GetToken())

	assert.NotNil(t, result)
}
