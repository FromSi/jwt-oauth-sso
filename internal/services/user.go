package services

//go:generate mockgen -destination=../../mocks/services/mock_query_user.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services QueryUserService
type QueryUserService interface {
	GenerateUUID() string
	HashPassword(string) (string, error)
	CheckHashedPasswordAndNativePassword(string, string) error
}

//go:generate mockgen -destination=../../mocks/services/mock_mutable_user.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services MutableUserService
type MutableUserService interface {
	CreateUserByUUIDAndEmailAndHashedPassword(string, string, string) error
	UpdatePasswordByUUIDAndHashedPassword(string, string) error
}

//go:generate mockgen -destination=../../mocks/services/mock_user.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services UserService
type UserService interface {
	QueryUserService
	MutableUserService
}
