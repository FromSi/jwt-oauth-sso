package repositories

import "gorm.io/gorm"

type GormUser struct {
	UUID      string `gorm:"unique;not null"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	CreatedAt int    `gorm:"not null"`
	UpdatedAt int    `gorm:"not null"`
}

func (receiver *GormUser) TableName() string {
	return "users"
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
	db          *gorm.DB
	userBuilder UserBuilder
}

func NewGormUserRepository(
	db *gorm.DB,
	userBuilder UserBuilder,
) (*GormUserRepository, error) {
	err := db.AutoMigrate(&GormUser{})

	if err != nil {
		return nil, err
	}

	return &GormUserRepository{
		db:          db,
		userBuilder: userBuilder,
	}, nil
}

func (receiver *GormUserRepository) GetUserByEmail(email string) User {
	var gormUser GormUser

	result := receiver.
		db.
		Model(&GormUser{}).
		First(&gormUser, &GormUser{Email: email})

	if result.RowsAffected == 0 {
		return nil
	}

	return &gormUser
}

func (receiver *GormUserRepository) GetUserByUUID(uuid string) User {
	var gormUser GormUser

	result := receiver.
		db.
		Model(&GormUser{}).
		First(&gormUser, &GormUser{UUID: uuid})

	if result.RowsAffected == 0 {
		return nil
	}

	return &gormUser
}

func (receiver *GormUserRepository) CreateUser(user User) error {
	gormUser, err := receiver.
		userBuilder.
		NewFromUser(user).
		BuildToGorm()

	if err != nil {
		return err
	}

	return receiver.
		db.
		Model(&GormUser{}).
		Create(gormUser).
		Error
}

func (receiver *GormUserRepository) UpdatePasswordByUUIDAndPasswordAndUpdatedAt(
	uuid string,
	password string,
	updatedAt int,
) error {
	return receiver.
		db.
		Model(&GormUser{}).
		Where(&GormUser{UUID: uuid}).
		UpdateColumns(&GormUser{Password: password, UpdatedAt: updatedAt}).
		Error
}
