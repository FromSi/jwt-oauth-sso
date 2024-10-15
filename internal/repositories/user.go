package repositories

type QueryUserRepository interface {
	HasUserByUUID(string) bool
	HasUserByEmail(string) bool
	GetUserByEmailAndPassword(string, string) User
}

type MutableUserRepository interface {
	CreateUser(User) error
	UpdatePassword(uuid string, password string, updatedAt int) error
}

type UserRepository interface {
	QueryUserRepository
	MutableUserRepository
}

type QueryUser interface {
	GetUUID() string
	GetEmail() string
	GetPassword() string
	GetCreatedAt() int
	GetUpdatedAt() int
}

type MutableUser interface {
	SetUUID(string)
	SetEmail(string)
	SetPassword(string)
	SetCreatedAt(int)
	SetUpdatedAt(int)
}

type User interface {
	QueryUser
	MutableUser
}
