package api

import (
	"fmt"
	"net/http"

	"github.com/forgoty/go-todo/pkg/web"
)

func (hs *HTTPServer) registerRoutes() {
	r := hs.routeRegister
	m := hs.web

	m.Use(web.MiddlewareLogger())
	m.Use(web.MiddlewareRecover())

	handle := func(ctx web.Context) error {
		fmt.Println(ctx.Path())
		return ctx.String(http.StatusOK, "Hello world\n")
	}
	r.Get("/api", handle)

	//User
	hs.addUserRoutes()

	hs.routeRegister.Register(hs.web.Router())
}
