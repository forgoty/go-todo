package http

import (
	"fmt"
	"net/http"

	"github.com/forgoty/go-todo/pkg/api/routing"
	"github.com/forgoty/go-todo/pkg/web"
)

func AddRoutes(r routing.RouteRegister) {
	r.Post("/login", login)
}

func login(ctx web.Context) error {
	fmt.Println(ctx.Path())
	return ctx.String(http.StatusOK, "Hello login\n")
}
