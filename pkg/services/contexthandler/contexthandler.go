package contexthandler

import (
	"github.com/forgoty/go-todo/pkg/models"
	"github.com/forgoty/go-todo/pkg/web"
)

type ContextHandler struct {
	//AuthService
}

func (h *ContextHandler) Middleware(next web.HandlerFunc) web.HandlerFunc {
	return func(c web.Context) error {
		reqContext := &models.ReqContext{
			Context: c,
		}

		switch {
		case h.initContextWithAnonymousUser(reqContext):
		}
		return next(reqContext)
	}
}

func (h *ContextHandler) initContextWithAnonymousUser(reqContext *models.ReqContext) bool {
	reqContext.IsSignedIn = false
	reqContext.AllowAnonymous = true
	reqContext.SignedInUser = &models.SignedInUser{IsAnonymous: true}
	return true
}
