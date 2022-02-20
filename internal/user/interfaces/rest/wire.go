// +build wireinject

package rest

import (
	"time"

	"github.com/forgoty/go-todo/internal/user/app"
	"github.com/forgoty/go-todo/internal/user/service/auth"
	"github.com/forgoty/go-todo/internal/user/service/contexthandler"
	"github.com/google/wire"
)

func provideUserController(salt auth.Salt, signinKey auth.SignInKey, tokenTTL time.Duration) (*userController, error) {
	wire.Build(
		wire.Struct(new(userController), "*"),
		app.NewUserApplication,
		auth.NewAuthService,
		contexthandler.NewContextHandler,
	)
	return &userController{}, nil
}
