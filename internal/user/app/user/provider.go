package user

import (
	"github.com/forgoty/go-todo/internal/user/app/user/getuser"
	"github.com/google/wire"
)

var ProvideUserSet wire.ProviderSet = wire.NewSet(
	getuser.ProvideGetUserQueryHandler,
)
