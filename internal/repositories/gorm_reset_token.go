package repositories

import (
	"gorm.io/gorm"
	"time"
)

type GormResetToken struct {
	Token     string `gorm:"not null;uniqueIndex:idx_token_useruuid"`
	UserUUID  string `gorm:"not null;uniqueIndex:idx_token_useruuid"`
	ExpiresAt int    `gorm:"not null"`
	CreatedAt int    `gorm:"not null"`
}

func (receiver *GormResetToken) TableName() string {
	return "reset_tokens"
}

func (receiver *GormResetToken) GetToken() string {
	return receiver.Token
}

func (receiver *GormResetToken) GetUserUUID() string {
	return receiver.UserUUID
}

func (receiver *GormResetToken) GetExpiresAt() int {
	return receiver.ExpiresAt
}

func (receiver *GormResetToken) GetCreatedAt() int {
	return receiver.CreatedAt
}

func (receiver *GormResetToken) SetToken(value string) {
	receiver.Token = value
}

func (receiver *GormResetToken) SetUserUUID(value string) {
	receiver.UserUUID = value
}

func (receiver *GormResetToken) SetExpiresAt(value int) {
	receiver.ExpiresAt = value
}

func (receiver *GormResetToken) SetCreatedAt(value int) {
	receiver.CreatedAt = value
}

type GormResetTokenRepository struct {
	db                *gorm.DB
	resetTokenBuilder ResetTokenBuilder
}

func NewGormResetTokenRepository(
	db *gorm.DB,
	resetTokenBuilder ResetTokenBuilder,
) (*GormResetTokenRepository, error) {
	err := db.AutoMigrate(&GormResetToken{})

	if err != nil {
		return nil, err
	}

	return &GormResetTokenRepository{
		db:                db,
		resetTokenBuilder: resetTokenBuilder,
	}, nil
}

func (receiver *GormResetTokenRepository) GetActiveResetTokenByToken(token string) ResetToken {
	var gormResetToken GormResetToken

	result := receiver.db.
		Model(&GormResetToken{}).
		Where("token = ? AND expires_at > ?", token, int(time.Now().Unix())).
		First(&gormResetToken)

	if result.RowsAffected == 0 {
		return nil
	}

	return &gormResetToken
}

func (receiver *GormResetTokenRepository) CreateResetToken(resetToken ResetToken) error {
	gormResetToken, err := receiver.
		resetTokenBuilder.
		NewFromResetToken(resetToken).
		BuildToGorm()

	if err != nil {
		return err
	}

	return receiver.
		db.
		Model(&GormResetToken{}).
		Create(gormResetToken).
		Error
}

func (receiver *GormResetTokenRepository) DeleteResetToken(token string) error {
	return receiver.
		db.
		Delete(&GormResetToken{}, &GormResetToken{Token: token}).
		Error
}
