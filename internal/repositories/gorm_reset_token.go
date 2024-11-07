package repositories

import (
	"gorm.io/gorm"
)

const (
	GormResetTokenTokenDefault     = ""
	GormResetTokenUserUUIDDefault  = ""
	GormResetTokenExpiredAtDefault = 1
	GormResetTokenCreatedAtDefault = 1
)

type GormResetToken struct {
	Token     string `gorm:"not null;uniqueIndex:idx_token_useruuid"`
	UserUUID  string `gorm:"not null;uniqueIndex:idx_token_useruuid"`
	ExpiredAt int    `gorm:"not null"`
	CreatedAt int    `gorm:"not null"`
}

func NewGormResetToken() *GormResetToken {
	return &GormResetToken{
		Token:     GormResetTokenTokenDefault,
		UserUUID:  GormResetTokenUserUUIDDefault,
		ExpiredAt: GormResetTokenExpiredAtDefault,
		CreatedAt: GormResetTokenCreatedAtDefault,
	}
}

func NewGormResetTokenByResetToken(resetToken ResetToken) *GormResetToken {
	return &GormResetToken{
		Token:     resetToken.GetToken(),
		UserUUID:  resetToken.GetUserUUID(),
		ExpiredAt: resetToken.GetExpiredAt(),
		CreatedAt: resetToken.GetCreatedAt(),
	}
}

func (receiver *GormResetToken) GetToken() string {
	return receiver.Token
}

func (receiver *GormResetToken) GetUserUUID() string {
	return receiver.UserUUID
}

func (receiver *GormResetToken) GetExpiredAt() int {
	return receiver.ExpiredAt
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

func (receiver *GormResetToken) SetExpiredAt(value int) {
	receiver.ExpiredAt = value
}

func (receiver *GormResetToken) SetCreatedAt(value int) {
	receiver.CreatedAt = value
}

type GormResetTokenRepository struct {
	db *gorm.DB
}

func NewGormResetTokenRepository(db *gorm.DB) (*GormResetTokenRepository, error) {
	err := db.AutoMigrate(&GormResetToken{})

	if err != nil {
		return nil, err
	}

	return &GormResetTokenRepository{db: db}, nil
}

func (receiver *GormResetTokenRepository) HasToken(token string) bool {
	var exists bool

	receiver.
		db.
		Model(&GormResetToken{}).
		Select("count(*) > 0").
		Find(&exists, &GormResetToken{Token: token})

	return exists
}

func (receiver *GormResetTokenRepository) GetResetTokenByToken(token string) ResetToken {
	var gormResetToken GormResetToken

	result := receiver.
		db.
		Model(&GormResetToken{}).
		First(&gormResetToken, &GormResetToken{Token: token})

	if result.RowsAffected == 0 {
		return nil
	}

	return &gormResetToken
}

func (receiver *GormResetTokenRepository) CreateResetToken(resetToken ResetToken) error {
	gormResetToken := NewGormResetTokenByResetToken(resetToken)

	return receiver.
		db.
		Model(&GormResetToken{}).
		Create(NewGormResetTokenByResetToken(gormResetToken)).
		Error
}

func (receiver *GormResetTokenRepository) DeleteResetToken(token string) error {
	return receiver.
		db.
		Delete(&GormResetToken{}, &GormResetToken{Token: token}).
		Error
}
