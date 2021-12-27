package user

type IUserRepository interface {
	Create(User) error
	FindOneById(id string) (*User, error)
	FindOneByUsernameAndPassword(username, password string) (*User, error)
}
