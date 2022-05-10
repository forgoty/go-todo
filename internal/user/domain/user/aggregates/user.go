package aggregates

import (
	"github.com/forgoty/go-todo/internal/user/domain/user/commands"
	"github.com/forgoty/go-todo/internal/user/domain/user/valueobjects"
	"github.com/forgoty/go-todo/pkg/errorcollector"
)

type User struct {
	Id           string
	Username     valueobjects.Username
	PasswordHash string
}

func NewUser(c commands.RegisterUserCommand, pswHash string) (*User, error) {
	collector := make(errorcollector.Fields)
	_, err := valueobjects.NewPassword(c.Password)
	if err != nil {
		collector.Add("password", err)
	}
	username, err := valueobjects.NewUsername(c.Username)
	if err != nil {
		collector.Add("username", err)
	}
	if err = collector.Err(); err != nil {
		return nil, err
	}
	return &User{
		Id:           c.Id,
		Username:     username,
		PasswordHash: pswHash,
	}, nil
}
