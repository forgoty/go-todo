package command

import (
	"github.com/google/wire"
)

var ProvideCommandsSet wire.ProviderSet = wire.NewSet(
	ProvideLoginUserWithJWTCommandHandler,
	ProvideRegisterUserCommandHandler,
)
