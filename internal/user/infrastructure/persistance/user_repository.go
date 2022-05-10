package persistance

import (
	"errors"

	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type InMemoryUserRepository struct {
	users map[string]*aggregates.User
	l     logger.Logger
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*aggregates.User),
		l:     logger.New("inmemoryuserrepo"),
	}
}

func (r *InMemoryUserRepository) FindOneById(id string) (*aggregates.User, error) {
	r.l.Info("Users DB:", r.users)
	u, ok := r.users[id]
	if !ok {
		return nil, errors.New("Not Found!")
	}
	return u, nil
}

func (r *InMemoryUserRepository) Create(u aggregates.User) error {
	r.users[u.Id] = &u
	r.l.Info("Users DB:", r.users)
	return nil
}

func (r *InMemoryUserRepository) FindOneByUsernameAndPassword(username, password string) (*aggregates.User, error) {
	r.l.Info("Users DB:", r.users)
	for id := range r.users {
		if r.users[id].Username.Equals(username) && r.users[id].PasswordHash == password {
			return r.users[id], nil
		}
	}
	return nil, errors.New("Not Found!")
}
