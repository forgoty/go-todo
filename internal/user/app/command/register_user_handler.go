package command

import (
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/internal/user/domain/user/commands"
	"github.com/forgoty/go-todo/internal/user/service/register"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type RegisterUserCommandHandler struct {
	log             logger.Logger
	userRepo        aggregates.IUserRepository
	registerService *register.RegisterService
}

func ProvideRegisterUserCommandHandler(ur aggregates.IUserRepository, regs *register.RegisterService) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{
		log:             logger.New("RegisterUserCommandHandler"),
		userRepo:        ur,
		registerService: regs,
	}
}

func (h *RegisterUserCommandHandler) Handle(c commands.RegisterUserCommand) error {
	return h.registerService.RegisterUser(c)
}
