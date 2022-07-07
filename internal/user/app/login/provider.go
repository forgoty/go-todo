package login

import (
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
	"github.com/google/wire"
)

func ProvideLoginUserWithJWTCommandHandler(ur aggregates.IUserRepository, ls *LoginService) *LoginUserWithJWTCommandHandler {
	return &LoginUserWithJWTCommandHandler{
		log:          logger.New("LoginUserWithJWTCommandHandler"),
		userRepo:     ur,
		loginService: ls,
	}
}

var ProvideLoginSet wire.ProviderSet = wire.NewSet(
	ProvideLoginUserWithJWTCommandHandler,
)
