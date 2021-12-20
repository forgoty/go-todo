package user

type IUserRepository interface {
	FindOneById(id string) *User
}
