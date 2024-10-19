package services

type QueryUserService interface {
	GenerateUUID() string
	HashPassword(string) (string, error)
	CheckPasswordByHashAndPassword(string, string) error
}

type MutableUserService interface {
}

type UserService interface {
	QueryUserService
	MutableUserService
}
