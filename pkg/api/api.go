package api

import (
	"fmt"
	"net/http"
	"github.com/swaggo/echo-swagger"

	_ "github.com/forgoty/go-todo/docs"
	user "github.com/forgoty/go-todo/internal/user/interfaces/rest"
	"github.com/forgoty/go-todo/pkg/api/routing"
	"github.com/forgoty/go-todo/pkg/web"
)

func (hs *HTTPServer) registerRoutes() {
	r := hs.routeRegister
	m := hs.web

	m.Use(web.MiddlewareLogger())
	m.Use(web.MiddlewareRecover())

	r.Get("/swagger/*", echoSwagger.WrapHandler)

	hello := func(ctx web.Context) error {
		fmt.Println(ctx.Path())
		return ctx.String(http.StatusOK, "Hello world\n")
	}
	r.Group("/api/v1", func(rr routing.RouteRegister) {
		rr.Get("/hello", hello)

		//User
		user.RegisterRoutesAndMiddlewares(rr, m)
	})

	hs.routeRegister.Register(hs.web.Router())
}
