package models

import (
	"github.com/forgoty/go-todo/pkg/web"
)

// Wrapper around echo context with custom features
type ReqContext struct {
	*SignedInUser
	web.Context
	IsSignedIn     bool
	IsRenderCall   bool
	AllowAnonymous bool
	SkipCache      bool
	LookupTokenErr error
}
