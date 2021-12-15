package api

import (
	"fmt"
	"net/http"

	user_public_http "github.com/forgoty/go-todo/internal/user/interfaces/public/http"
	"github.com/forgoty/go-todo/pkg/web"
)

func (hs *HTTPServer) registerRoutes() {
	r := hs.RouteRegister
	handle := func(ctx web.Context) error {
		fmt.Println(ctx.Path())
		return ctx.String(http.StatusOK, "Hello world\n")
	}
	r.Get("/api", handle)
	user_public_http.AddRoutes(r)
}
