package command

import (
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/internal/user/domain/user/commands"
	"github.com/forgoty/go-todo/internal/user/service/register"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type LoginUserWithJWTCommandHandler struct {
	log             logger.Logger
	userRepo        aggregates.IUserRepository
	registerService *register.RegisterService
}

func ProvideLoginUserWithJWTCommandHandler(ur aggregates.IUserRepository, regs *register.RegisterService) *LoginUserWithJWTCommandHandler {
	return &LoginUserWithJWTCommandHandler{
		log:             logger.New("LoginUserWithJWTCommandHandler"),
		userRepo:        ur,
		registerService: regs,
	}
}

func (h *LoginUserWithJWTCommandHandler) Handle(c commands.LoginUserWithJWTCommand) (string, error) {
	return "", nil
}
