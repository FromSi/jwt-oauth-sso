package repositories

//go:generate mockgen -destination=../mocks/repositories/mock_query_reset_token_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories QueryResetTokenRepository
type QueryResetTokenRepository interface {
	HasToken(string) bool
	GetActiveResetTokenByToken(string) ResetToken
}

//go:generate mockgen -destination=../mocks/repositories/mock_mutable_reset_token_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories MutableResetTokenRepository
type MutableResetTokenRepository interface {
	CreateResetToken(ResetToken) error
	DeleteResetToken(string) error
}

//go:generate mockgen -destination=../mocks/repositories/mock_reset_token_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories ResetTokenRepository
type ResetTokenRepository interface {
	QueryResetTokenRepository
	MutableResetTokenRepository
}

//go:generate mockgen -destination=../mocks/repositories/mock_query_reset_token.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories QueryResetToken
type QueryResetToken interface {
	GetToken() string
	GetUserUUID() string
	GetExpiresAt() int
	GetCreatedAt() int
}

//go:generate mockgen -destination=../mocks/repositories/mock_mutable_reset_token.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories MutableResetToken
type MutableResetToken interface {
	SetToken(string)
	SetUserUUID(string)
	SetExpiresAt(int)
	SetCreatedAt(int)
}

//go:generate mockgen -destination=../mocks/repositories/mock_reset_token.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories ResetToken
type ResetToken interface {
	QueryResetToken
	MutableResetToken
}
