package handler

import (
	"github.com/forgoty/go-todo/internal/user/app/command"
	"github.com/forgoty/go-todo/internal/user/domain/user"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type RegisterUserHandler struct {
	log      logger.Logger
	userRepo user.IUserRepository
}

func NewRegisterUserHandler(repo user.IUserRepository, log logger.Logger) *RegisterUserHandler {
	return &RegisterUserHandler{
		userRepo: repo,
		log:      log,
	}
}

func (h *RegisterUserHandler) Handle(c command.RegisterUserCommand) error {
	h.log.Info("Register " + c.Username)
	return nil
}
