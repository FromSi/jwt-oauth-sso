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
	assert.Equal(t, gormResetToken.ExpiresAt, GormResetTokenExpiredAtDefault)
	assert.Equal(t, gormResetToken.CreatedAt, GormResetTokenCreatedAtDefault)
}

func Test_NewGormResetTokenByResetToken(t *testing.T) {
	gormResetTokenTemp := NewGormResetToken()
	gormResetToken := NewGormResetTokenByResetToken(gormResetTokenTemp)

	assert.NotNil(t, gormResetToken)

	assert.Equal(t, gormResetToken.Token, GormResetTokenTokenDefault)
	assert.Equal(t, gormResetToken.UserUUID, GormResetTokenUserUUIDDefault)
	assert.Equal(t, gormResetToken.ExpiresAt, GormResetTokenExpiredAtDefault)
	assert.Equal(t, gormResetToken.CreatedAt, GormResetTokenCreatedAtDefault)

}

func TestGormResetToken_GetToken(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.Token = "1"

	assert.Equal(t, gormResetToken.GetToken(), "1")
}

func TestGormResetToken_GetUserUUID(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.UserUUID = "1"

	assert.Equal(t, gormResetToken.GetUserUUID(), "1")
}

func TestGormResetToken_GetExpiredAt(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.ExpiresAt = 1

	assert.Equal(t, gormResetToken.GetExpiresAt(), 1)
}

func TestGormResetToken_GetCreatedAt(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.CreatedAt = 1

	assert.Equal(t, gormResetToken.GetCreatedAt(), 1)
}

func TestGormResetToken_SetToken(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.SetToken("1")

	assert.Equal(t, gormResetToken.Token, "1")
}

func TestGormResetToken_SetUserUUID(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.SetUserUUID("1")

	assert.Equal(t, gormResetToken.UserUUID, "1")
}

func TestGormResetToken_SetExpiredAt(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.SetExpiresAt(1)

	assert.Equal(t, gormResetToken.ExpiresAt, 1)
}

func TestGormResetToken_SetCreatedAt(t *testing.T) {
	gormResetToken := NewGormResetToken()

	gormResetToken.SetCreatedAt(1)

	assert.Equal(t, gormResetToken.CreatedAt, 1)
}

func Test_NewGormResetTokenRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	assert.Nil(t, err)
	assert.NotNil(t, db)

	gormResetTokenRepository, err := NewGormResetTokenRepository(db)

	assert.Nil(t, err)
	assert.NotNil(t, gormResetTokenRepository)

	stmt := &gorm.Statement{DB: db}
	err = stmt.Parse(&GormResetToken{})

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

func TestGormResetTokenRepository_CreateToken_And_HasToken(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormResetTokenRepository, _ := NewGormResetTokenRepository(db)

	gormResetToken := NewGormResetToken()

	gormResetToken.SetToken("1")

	err := gormResetTokenRepository.CreateResetToken(gormResetToken)

	assert.Nil(t, err)

	exists := gormResetTokenRepository.HasToken(gormResetToken.GetToken())

	assert.True(t, exists)

	exists = gormResetTokenRepository.HasToken("0")

	assert.False(t, exists)
}

func TestGormResetTokenRepository_CreateToken_And_GetActiveResetTokenByToken(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormResetTokenRepository, _ := NewGormResetTokenRepository(db)

	gormResetToken := NewGormResetToken()

	gormResetToken.SetToken("1")
	gormResetToken.SetUserUUID("2")
	gormResetToken.SetExpiresAt(int(time.Now().AddDate(0, 0, 1).Unix()))

	err := gormResetTokenRepository.CreateResetToken(gormResetToken)

	assert.Nil(t, err)

	resetTokenByToken := gormResetTokenRepository.
		GetActiveResetTokenByToken(gormResetToken.GetToken())

	assert.NotNil(t, resetTokenByToken)

	assert.Equal(t, resetTokenByToken.GetToken(), gormResetToken.GetToken())
	assert.Equal(t, resetTokenByToken.GetUserUUID(), gormResetToken.GetUserUUID())
	assert.Equal(t, resetTokenByToken.GetExpiresAt(), gormResetToken.GetExpiresAt())

	resetTokenByToken = gormResetTokenRepository.GetActiveResetTokenByToken("5")

	assert.Nil(t, resetTokenByToken)

	gormResetToken = NewGormResetToken()

	gormResetToken.SetToken("2")
	gormResetToken.SetUserUUID("2")
	gormResetToken.SetExpiresAt(0)

	err = gormResetTokenRepository.CreateResetToken(gormResetToken)

	assert.Nil(t, err)

	assert.Nil(t, resetTokenByToken)
}

func TestGormResetTokenRepository_CreateToken_And_DeleteResetToken(t *testing.T) {
	db, _ := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	gormResetTokenRepository, _ := NewGormResetTokenRepository(db)

	gormResetToken := NewGormResetToken()

	gormResetToken.SetToken("1")

	err := gormResetTokenRepository.CreateResetToken(gormResetToken)

	assert.Nil(t, err)

	result := gormResetTokenRepository.HasToken(gormResetToken.GetToken())

	assert.True(t, result)

	err = gormResetTokenRepository.DeleteResetToken(gormResetToken.GetToken())

	assert.Nil(t, err)

	result = gormResetTokenRepository.HasToken(gormResetToken.GetToken())

	assert.False(t, result)
}
