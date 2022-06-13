// +build wireinject

package rest

import (
	"time"

	"github.com/forgoty/go-todo/internal/user/app"
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/internal/user/infrastructure/persistence"
	"github.com/forgoty/go-todo/internal/user/service/auth"
	"github.com/forgoty/go-todo/internal/user/service/contexthandler"
	"github.com/forgoty/go-todo/internal/user/service/login"
	"github.com/forgoty/go-todo/internal/user/service/register"
	"github.com/google/wire"
)

func provideUserController(salt auth.Salt, signinKey auth.SignInKey, tokenTTL time.Duration) (*userController, error) {
	wire.Build(
		persistence.NewInMemoryUserRepository,
		app.NewUserApplication,
		auth.NewAuthService,
		login.NewLoginService,
		register.NewRegisterService,
		contexthandler.NewContextHandler,
		wire.Bind(new(aggregates.IUserRepository), new(*persistence.InMemoryUserRepository)),
		wire.Struct(new(userController), "*"),
	)
	return &userController{}, nil
}
