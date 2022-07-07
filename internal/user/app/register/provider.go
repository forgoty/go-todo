package register

import (
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
	"github.com/google/wire"
)

func ProvideRegisterUserCommandHandler(ur aggregates.IUserRepository, regs *RegisterService) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{
		log:             logger.New("RegisterUserCommandHandler"),
		userRepo:        ur,
		registerService: regs,
	}
}

var ProvideRegisterSet wire.ProviderSet = wire.NewSet(
	ProvideRegisterUserCommandHandler,
)
