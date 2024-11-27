package repositories

import "errors"

type BaseUserBuilder struct {
	user GormUser
}

func NewBaseUserBuilder() *BaseUserBuilder {
	return &BaseUserBuilder{
		user: GormUser{},
	}
}

func (receiver *BaseUserBuilder) New() UserBuilder {
	return &BaseUserBuilder{
		user: GormUser{},
	}
}

func (receiver *BaseUserBuilder) NewFromUser(user User) UserBuilder {
	return receiver.
		New().
		SetUUID(user.GetUUID()).
		SetEmail(user.GetEmail()).
		SetPassword(user.GetPassword()).
		SetCreatedAt(user.GetCreatedAt()).
		SetUpdatedAt(user.GetUpdatedAt())
}

func (receiver *BaseUserBuilder) Build() (User, error) {
	return receiver.BuildToGorm()
}

func (receiver *BaseUserBuilder) BuildToGorm() (*GormUser, error) {
	if len(receiver.user.GetUUID()) == 0 {
		return nil, errors.New("uuid must not be empty")
	}

	return &receiver.user, nil
}

func (receiver *BaseUserBuilder) SetUUID(value string) UserBuilder {
	receiver.user.SetUUID(value)

	return receiver
}

func (receiver *BaseUserBuilder) SetEmail(value string) UserBuilder {
	receiver.user.SetEmail(value)

	return receiver
}

func (receiver *BaseUserBuilder) SetPassword(value string) UserBuilder {
	receiver.user.SetPassword(value)

	return receiver
}

func (receiver *BaseUserBuilder) SetCreatedAt(value int) UserBuilder {
	receiver.user.SetCreatedAt(value)

	return receiver
}

func (receiver *BaseUserBuilder) SetUpdatedAt(value int) UserBuilder {
	receiver.user.SetUpdatedAt(value)

	return receiver
}
