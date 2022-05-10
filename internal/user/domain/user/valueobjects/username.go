package valueobjects

import (
	"github.com/forgoty/go-todo/internal/user/domain/user/valueobjects/validators"
)

type Username string

func NewUsername(username string) (Username, error) {
	if err := validators.ValidateUsername(username); err != nil {
		return Username(""), err
	}
	return Username(username), nil
}

func (u Username) Equals(other string) bool {
	return string(u) == other
}
