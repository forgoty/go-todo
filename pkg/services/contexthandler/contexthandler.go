package contexthandler

import (
	"fmt"

	"github.com/forgoty/go-todo/pkg/models"
	"github.com/forgoty/go-todo/pkg/web"
)

type ContextHandler struct {
	// Logger
}

func (h *ContextHandler) Middleware(next web.HandlerFunc) web.HandlerFunc {
	return func(c web.Context) error {
		reqContext := &models.ReqContext{
			Context: c,
		}
		fmt.Println(reqContext.Scheme())
		// switch {
		// case: auth
		// }
		return next(reqContext)
	}
}

func (h *ContextHandler) MiddlewareH(next web.HandlerFunc) web.HandlerFunc {
	return func(c web.Context) error {
		fmt.Println(c)
		// switch {
		// case: auth
		// }
		return next(c)
	}
}
