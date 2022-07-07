// +build wireinject

package rest

import (
	"time"

	"github.com/forgoty/go-todo/internal/user/app"
	"github.com/forgoty/go-todo/internal/user/app/login"
	"github.com/forgoty/go-todo/internal/user/app/register"
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/internal/user/service/auth"
	"github.com/forgoty/go-todo/internal/user/service/contexthandler"
	"github.com/google/wire"
)

func provideUserController(salt auth.Salt, signinKey auth.SignInKey, tokenTTL time.Duration, userRepo aggregates.IUserRepository) (*userController, error) {
	wire.Build(
		app.NewUserApplication,
		wire.Struct(new(auth.PasswordManager), "*"),
		wire.Struct(new(auth.JWTService), "*"),
		login.NewLoginService,
		register.NewRegisterService,
		contexthandler.NewContextHandler,
		wire.Struct(new(userController), "*"),
	)
	return &userController{}, nil
}
