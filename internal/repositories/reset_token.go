package repositories

type QueryResetTokenRepository interface {
	HasToken(string) bool
	GetResetTokenByToken(string) ResetToken
}

type MutableResetTokenRepository interface {
	CreateResetToken(ResetToken) error
	DeleteResetToken(string) error
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
