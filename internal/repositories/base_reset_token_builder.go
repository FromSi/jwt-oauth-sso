package repositories

import "errors"

type BaseResetTokenBuilder struct {
	resetToken GormResetToken
}

func NewBaseResetTokenBuilder() *BaseResetTokenBuilder {
	return &BaseResetTokenBuilder{
		resetToken: GormResetToken{},
	}
}

func (receiver *BaseResetTokenBuilder) New() ResetTokenBuilder {
	return &BaseResetTokenBuilder{
		resetToken: GormResetToken{},
	}
}

func (receiver *BaseResetTokenBuilder) NewFromResetToken(resetToken ResetToken) ResetTokenBuilder {
	return receiver.
		New().
		SetToken(resetToken.GetToken()).
		SetUserUUID(resetToken.GetUserUUID()).
		SetExpiresAt(resetToken.GetExpiresAt()).
		SetCreatedAt(resetToken.GetCreatedAt())
}

func (receiver *BaseResetTokenBuilder) Build() (ResetToken, error) {
	return receiver.BuildToGorm()
}

func (receiver *BaseResetTokenBuilder) BuildToGorm() (*GormResetToken, error) {
	if len(receiver.resetToken.GetToken()) == 0 {
		return nil, errors.New("token must not be empty")
	}

	return &receiver.resetToken, nil
}

func (receiver *BaseResetTokenBuilder) SetToken(value string) ResetTokenBuilder {
	receiver.resetToken.SetToken(value)

	return receiver
}

func (receiver *BaseResetTokenBuilder) SetUserUUID(value string) ResetTokenBuilder {
	receiver.resetToken.SetUserUUID(value)

	return receiver
}

func (receiver *BaseResetTokenBuilder) SetExpiresAt(value int) ResetTokenBuilder {
	receiver.resetToken.SetExpiresAt(value)

	return receiver
}

func (receiver *BaseResetTokenBuilder) SetCreatedAt(value int) ResetTokenBuilder {
	receiver.resetToken.SetCreatedAt(value)

	return receiver
}
