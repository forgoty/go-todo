package query

import (
	"github.com/google/wire"
)

var ProvideQueriesSet wire.ProviderSet = wire.NewSet(
	ProvideFindUserBySigninQueryHandler,
	ProvideGetUserQueryHandler,
)
