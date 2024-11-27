package repositories

//go:generate mockgen -destination=../../mocks/repositories/mock_user_builder.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories UserBuilder
type UserBuilder interface {
	New() UserBuilder
	NewFromUser(User) UserBuilder
	Build() (User, error)
	BuildToGorm() (*GormUser, error)
	SetUUID(string) UserBuilder
	SetEmail(string) UserBuilder
	SetPassword(string) UserBuilder
	SetCreatedAt(int) UserBuilder
	SetUpdatedAt(int) UserBuilder
}
