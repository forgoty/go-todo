package app

import (
	"github.com/forgoty/go-todo/internal/user/app/login"
	"github.com/forgoty/go-todo/internal/user/app/register"
	"github.com/forgoty/go-todo/internal/user/app/user"
	"github.com/google/wire"
)

var applicationProviderSet wire.ProviderSet = wire.NewSet(
	login.ProvideLoginSet,
	register.ProvideRegisterSet,
	user.ProvideUserSet,
	wire.Struct(new(Commands), "*"),
	wire.Struct(new(Queries), "*"),
	wire.Struct(new(Application), "*"),
)
