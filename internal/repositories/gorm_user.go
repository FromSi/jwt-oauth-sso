package repositories

import "gorm.io/gorm"

const (
	GormUserUUIDDefault      = ""
	GormUserEmailDefault     = ""
	GormUserPasswordDefault  = ""
	GormUserCreatedAtDefault = 0
	GormUserUpdatedAtDefault = 0
)

type GormUser struct {
	UUID      string `gorm:"unique;not null"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	CreatedAt int    `gorm:"not null"`
	UpdatedAt int    `gorm:"not null"`
}

func NewGormUser() *GormUser {
	return &GormUser{
		UUID:      GormUserUUIDDefault,
		Email:     GormUserEmailDefault,
		Password:  GormUserPasswordDefault,
		CreatedAt: GormUserCreatedAtDefault,
		UpdatedAt: GormUserUpdatedAtDefault,
	}
}

func NewGormUserByUser(user User) *GormUser {
	return &GormUser{
		UUID:      user.GetUUID(),
		Email:     user.GetEmail(),
		Password:  user.GetPassword(),
		CreatedAt: user.GetCreatedAt(),
		UpdatedAt: user.GetUpdatedAt(),
	}
}

func (receiver *GormUser) GetUUID() string {
	return receiver.UUID
}

func (receiver *GormUser) GetEmail() string {
	return receiver.Email
}

func (receiver *GormUser) GetPassword() string {
	return receiver.Password
}

func (receiver *GormUser) GetCreatedAt() int {
	return receiver.CreatedAt
}

func (receiver *GormUser) GetUpdatedAt() int {
	return receiver.UpdatedAt
}

func (receiver *GormUser) SetUUID(value string) {
	receiver.UUID = value
}

func (receiver *GormUser) SetEmail(value string) {
	receiver.Email = value
}

func (receiver *GormUser) SetPassword(value string) {
	receiver.Password = value
}

func (receiver *GormUser) SetCreatedAt(value int) {
	receiver.CreatedAt = value
}

func (receiver *GormUser) SetUpdatedAt(value int) {
	receiver.UpdatedAt = value
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) (*GormUserRepository, error) {
	err := db.AutoMigrate(&GormUser{})

	if err != nil {
		return nil, err
	}

	return &GormUserRepository{db: db}, nil
}

func (receiver *GormUserRepository) HasUserByUUID(uuid string) bool {
	var exists bool

	receiver.db.Model(&GormUser{}).Select("count(*) > 0").Find(&exists, &GormUser{UUID: uuid})

	return exists
}

func (receiver *GormUserRepository) HasUserByEmail(email string) bool {
	var exists bool

	receiver.db.Model(&GormUser{}).Select("count(*) > 0").Find(&exists, &GormUser{Email: email})

	return exists
}

func (receiver *GormUserRepository) HasUserByEmailAndPassword(email string, password string) bool {
	var exists bool

	receiver.db.Model(&GormUser{}).Select("count(*) > 0").Find(&exists, &GormUser{Email: email, Password: password})

	return exists
}

func (receiver *GormUserRepository) HasUserByUUIDAndPassword(uuid string, password string) bool {
	var exists bool

	receiver.db.Model(&GormUser{}).Select("count(*) > 0").Find(&exists, &GormUser{UUID: uuid, Password: password})

	return exists
}

func (receiver *GormUserRepository) GetUserByEmail(email string) User {
	var gormUser GormUser

	result := receiver.db.Model(&GormUser{}).First(&gormUser, &GormUser{Email: email})

	if result.RowsAffected == 0 {
		return nil
	}

	return &gormUser
}

func (receiver *GormUserRepository) CreateUser(user User) error {
	gormUser := NewGormUserByUser(user)

	return receiver.db.Model(&GormUser{}).Create(NewGormUserByUser(gormUser)).Error
}

func (receiver *GormUserRepository) UpdatePassword(uuid string, password string, updatedAt int) error {
	return receiver.db.Model(&GormUser{}).Where(&GormUser{UUID: uuid}).UpdateColumns(&GormUser{Password: password, UpdatedAt: updatedAt}).Error
}
