// +build wireinject

package rest

import (
	"time"

	"github.com/forgoty/go-todo/internal/user/app"
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/internal/user/service/auth"
	"github.com/forgoty/go-todo/internal/user/service/contexthandler"
	"github.com/forgoty/go-todo/internal/user/service/login"
	"github.com/forgoty/go-todo/internal/user/service/register"
	"github.com/google/wire"
)

func provideUserController(salt auth.Salt, signinKey auth.SignInKey, tokenTTL time.Duration, userRepo aggregates.IUserRepository) (*userController, error) {
	wire.Build(
		app.NewUserApplication,
		auth.NewAuthService,
		login.NewLoginService,
		register.NewRegisterService,
		contexthandler.NewContextHandler,
		wire.Struct(new(userController), "*"),
	)
	return &userController{}, nil
}
