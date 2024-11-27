package repositories

//go:generate mockgen -destination=../../mocks/repositories/mock_reset_token_builder.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories ResetTokenBuilder
type ResetTokenBuilder interface {
	New() ResetTokenBuilder
	NewFromResetToken(ResetToken) ResetTokenBuilder
	Build() (ResetToken, error)
	BuildToGorm() (*GormResetToken, error)
	SetToken(string) ResetTokenBuilder
	SetUserUUID(string) ResetTokenBuilder
	SetExpiresAt(int) ResetTokenBuilder
	SetCreatedAt(int) ResetTokenBuilder
}
