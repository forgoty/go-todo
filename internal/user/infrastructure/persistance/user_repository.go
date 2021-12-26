package persistance

import (
	"errors"

	"github.com/forgoty/go-todo/internal/user/domain/user"
)

type InMemoryUserRepository struct {
	users map[string]*user.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*user.User),
	}
}

func (r *InMemoryUserRepository) FindOneById(id string) (*user.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, errors.New("Not Found!")
	}
	return u, nil
}

func (r *InMemoryUserRepository) Create(u user.User) error {
	r.users[u.Id] = &u
	return nil
}
