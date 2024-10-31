package repositories

//go:generate mockgen -destination=../mocks/repositories/mock_query_user_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories QueryUserRepository
type QueryUserRepository interface {
	HasUserByUUID(string) bool
	HasUserByEmail(string) bool
	HasUserByEmailAndPassword(string, string) bool
	HasUserByUUIDAndPassword(string, string) bool
	GetUserByEmail(string) User
}

//go:generate mockgen -destination=../mocks/repositories/mock_mutable_user_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories MutableUserRepository
type MutableUserRepository interface {
	CreateUser(User) error
	UpdatePassword(uuid string, password string, updatedAt int) error
}

//go:generate mockgen -destination=../mocks/repositories/mock_user_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories UserRepository
type UserRepository interface {
	QueryUserRepository
	MutableUserRepository
}

//go:generate mockgen -destination=../mocks/repositories/mock_query_user.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories QueryUser
type QueryUser interface {
	GetUUID() string
	GetEmail() string
	GetPassword() string
	GetCreatedAt() int
	GetUpdatedAt() int
}

//go:generate mockgen -destination=../mocks/repositories/mock_mutable_user.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories MutableUser
type MutableUser interface {
	SetUUID(string)
	SetEmail(string)
	SetPassword(string)
	SetCreatedAt(int)
	SetUpdatedAt(int)
}

//go:generate mockgen -destination=../mocks/repositories/mock_user.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories User
type User interface {
	QueryUser
	MutableUser
}
