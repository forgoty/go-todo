package api

import (
	"fmt"
	"github.com/forgoty/go-todo/pkg/web"
	"net/http"
)

func (hs *HTTPServer) registerRoutes() {
	r := hs.RouteRegister
	handle := func(ctx web.Context) error {
		fmt.Println(ctx.Path())
		return ctx.String(http.StatusOK, "Hello world\n")
	}
	r.Get("/api", handle)
}
