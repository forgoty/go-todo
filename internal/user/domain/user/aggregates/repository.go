package aggregates

type IUserRepository interface {
	Create(User) error
	FindOneById(id string) (*User, error)
	FindOneByUsername(username string) (*User, error)
}
