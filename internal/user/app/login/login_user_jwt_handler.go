package login

import (
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type LoginUserWithJWTCommandHandler struct {
	log          logger.Logger
	userRepo     aggregates.IUserRepository
	loginService *LoginService
}

func (h *LoginUserWithJWTCommandHandler) Handle(c LoginUserWithJWTCommand) (string, error) {
	return h.loginService.LoginUserJWT(c)
}
