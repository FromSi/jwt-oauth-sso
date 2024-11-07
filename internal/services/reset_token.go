package services

//go:generate mockgen -destination=../mocks/services/mock_query_reset_token_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services QueryResetTokenService
type QueryResetTokenService interface {
	GenerateToken() string
}

//go:generate mockgen -destination=../mocks/services/mock_mutable_reset_token_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services MutableResetTokenService
type MutableResetTokenService interface {
	SendNewResetTokenByUserEmail(string) error
	ResetPasswordByTokenAndNewPassword(string, string) error
	ResetPasswordByUserUUIDAndOldPasswordAndNewPassword(string, string, string) error
}

//go:generate mockgen -destination=../mocks/services/mock_reset_token_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services ResetTokenService
type ResetTokenService interface {
	QueryResetTokenService
	MutableResetTokenService
}
