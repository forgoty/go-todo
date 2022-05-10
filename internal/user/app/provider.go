package app

import (
	"github.com/forgoty/go-todo/internal/user/app/command"
	"github.com/forgoty/go-todo/internal/user/app/query"
	"github.com/google/wire"
)

var applicationProviderSet wire.ProviderSet = wire.NewSet(
	command.ProvideCommandsSet,
	query.ProvideQueriesSet,
	wire.Struct(new(Commands), "*"),
	wire.Struct(new(Queries), "*"),
	wire.Struct(new(Application), "*"),
)
