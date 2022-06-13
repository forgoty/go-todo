package command

import (
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/internal/user/domain/user/commands"
	"github.com/forgoty/go-todo/internal/user/service/login"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type LoginUserWithJWTCommandHandler struct {
	log          logger.Logger
	userRepo     aggregates.IUserRepository
	loginService *login.LoginService
}

func ProvideLoginUserWithJWTCommandHandler(ur aggregates.IUserRepository, ls *login.LoginService) *LoginUserWithJWTCommandHandler {
	return &LoginUserWithJWTCommandHandler{
		log:          logger.New("LoginUserWithJWTCommandHandler"),
		userRepo:     ur,
		loginService: ls,
	}
}

func (h *LoginUserWithJWTCommandHandler) Handle(c commands.LoginUserWithJWTCommand) (string, error) {
	return h.loginService.LoginUserJWT(c)
}
