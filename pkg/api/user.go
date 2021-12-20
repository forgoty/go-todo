package api

import (
	"fmt"
	"net/http"

	"github.com/forgoty/go-todo/pkg/web"
)

func (hs *HTTPServer) addUserRoutes() {
	r := hs.RouteRegister

	r.Post("/login", hs.login)
}

func (hs *HTTPServer) login(ctx web.Context) error {
	fmt.Println(ctx.Path())
	return ctx.String(http.StatusOK, "Hello login\n")
}
