package user

type IUserRepository interface {
	Create(User) error
	FindOneById(id string) (*User, error)
}
