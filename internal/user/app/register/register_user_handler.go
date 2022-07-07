package register

import (
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type RegisterUserCommandHandler struct {
	log             logger.Logger
	userRepo        aggregates.IUserRepository
	registerService *RegisterService
}

func (h *RegisterUserCommandHandler) Handle(c RegisterUserCommand) error {
	return h.registerService.RegisterUser(c)
}
