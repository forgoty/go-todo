package command

import (
	"errors"

	"github.com/forgoty/go-todo/internal/user/domain/user"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type RegisterUserCommand struct {
	Id       string `json:"-"`
	Username string `json:"username"`
}

type RegisterUserCommandHandler struct {
	log      logger.Logger
	userRepo user.IUserRepository
}

func NewRegisterUserCommandHandler(repo user.IUserRepository, log logger.Logger) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{
		userRepo: repo,
		log:      log,
	}
}

func (h *RegisterUserCommandHandler) Handle(c RegisterUserCommand) error {
	user := user.User{
		Id:       c.Id,
		Username: c.Username,
	}
	if err := h.userRepo.Create(user); err != nil {
		return errors.New("Cannot Create user")
	}
	return nil
}
