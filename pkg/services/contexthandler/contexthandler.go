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
		reqContext := models.NewContext(c)
		return next(reqContext)
	}
}
