package repositories

type QueryResetTokenRepository interface {
	HasTokenByToken(string) bool
}

type MutableResetTokenRepository interface {
	CreateResetToken(ResetToken) error
	DeleteResetTokenByToken(string) error
}

type ResetTokenRepository interface {
	QueryResetTokenRepository
	MutableResetTokenRepository
}

type QueryResetToken interface {
	GetToken() string
	GetUserUUID() string
	GetExpiredAt() int
	GetCreatedAt() int
}

type MutableResetToken interface {
	SetToken(string)
	SetUserUUID(string)
	SetExpiredAt(int)
	SetCreatedAt(int)
}

type ResetToken interface {
	QueryResetToken
	MutableResetToken
}
