// +build wireinject

package app

import (
	"github.com/forgoty/go-todo/internal/user/app/login"
	"github.com/forgoty/go-todo/internal/user/app/register"
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/google/wire"
)

func NewUserApplication(userRepo aggregates.IUserRepository, regs *register.RegisterService, ls *login.LoginService) *Application {
	wire.Build(applicationProviderSet)
	return &Application{}
}
