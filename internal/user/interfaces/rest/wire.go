// +build wireinject

package rest

import (
	"time"

	"github.com/forgoty/go-todo/internal/user/app"
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/internal/user/infrastructure/persistance"
	"github.com/forgoty/go-todo/internal/user/service/auth"
	"github.com/forgoty/go-todo/internal/user/service/contexthandler"
	"github.com/forgoty/go-todo/internal/user/service/register"
	"github.com/google/wire"
)

func provideUserController(salt auth.Salt, signinKey auth.SignInKey, tokenTTL time.Duration) (*userController, error) {
	wire.Build(
		persistance.NewInMemoryUserRepository,
		app.NewUserApplication,
		auth.NewAuthService,
		register.NewRegisterService,
		contexthandler.NewContextHandler,
		wire.Bind(new(aggregates.IUserRepository), new(*persistance.InMemoryUserRepository)),
		wire.Struct(new(userController), "*"),
	)
	return &userController{}, nil
}
