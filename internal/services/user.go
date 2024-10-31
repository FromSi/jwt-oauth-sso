package services

//go:generate mockgen -destination=../mocks/services/mock_query_user_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services QueryUserService
type QueryUserService interface {
	GenerateUUID() string
	HashPassword(string) (string, error)
	CheckPasswordByHashAndPassword(string, string) error
}

//go:generate mockgen -destination=../mocks/services/mock_mutable_user_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services MutableUserService
type MutableUserService interface {
	CreateUserByUUIDAndEmailAndPassword(string, string, string) error
}

//go:generate mockgen -destination=../mocks/services/mock_user_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services UserService
type UserService interface {
	QueryUserService
	MutableUserService
}
