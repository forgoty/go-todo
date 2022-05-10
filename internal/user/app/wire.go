// +build wireinject

package app

import (
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/internal/user/service/register"
	"github.com/google/wire"
)

func NewUserApplication(userRepo aggregates.IUserRepository, regs *register.RegisterService) *Application {
	wire.Build(applicationProviderSet)
	return &Application{}
}
