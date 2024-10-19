package services

type QueryResetTokenService interface {
	GenerateToken() string
}

type MutableResetTokenService interface {
}

type ResetTokenService interface {
	QueryResetTokenService
	MutableResetTokenService
}
