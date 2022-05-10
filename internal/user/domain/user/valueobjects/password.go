package valueobjects

import (
	"errors"

	"github.com/forgoty/go-todo/internal/user/domain/user/valueobjects/validators"
)

type Password string

func NewPassword(password string) (Password, error) {
	if err := validators.ValidatePassword(password); err != nil {
		return Password(""), errors.New(err.Error())
	}
	return Password(password), nil
}
