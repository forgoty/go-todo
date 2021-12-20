package handler

import (
	"errors"

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

func (h *RegisterUserHandler) Handle() error { return errors.New("Not Implemented") }
