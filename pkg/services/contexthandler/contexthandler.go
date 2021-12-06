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
		reqContext.Logger().SetPrefix("context")
		reqContext.Logger().SetLevel(2) // INFO

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
