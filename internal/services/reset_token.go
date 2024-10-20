package services

type QueryResetTokenService interface {
	GenerateToken() string
}

type MutableResetTokenService interface {
	ResetPasswordByTokenAndNewPassword(string, string) error
	ResetPasswordByUserUUIDAndOldPasswordAndNewPassword(string, string, string) error
}

type ResetTokenService interface {
	QueryResetTokenService
	MutableResetTokenService
}
