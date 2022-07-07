package aggregates

import (
	"github.com/forgoty/go-todo/internal/user/domain/user/valueobjects"
	"github.com/forgoty/go-todo/pkg/errorcollector"
)

type User struct {
	Id           string
	Username     valueobjects.Username
	PasswordHash string
	UserProfile  valueobjects.UserProfile
}

func NewUser(id, username, password, pswHash string) (*User, error) {
	collector := make(errorcollector.Fields)
	_, err := valueobjects.NewPassword(password)
	if err != nil {
		collector.Add("password", err)
	}
	usernameVo, err := valueobjects.NewUsername(username)
	if err != nil {
		collector.Add("username", err)
	}
	if err = collector.Err(); err != nil {
		return nil, err
	}
	return &User{
		Id:           id,
		Username:     usernameVo,
		PasswordHash: pswHash,
	}, nil
}
