package api

import (
	"fmt"
	"net/http"

	"github.com/forgoty/go-todo/internal/user/app/command"
	"github.com/forgoty/go-todo/pkg/web"
)

func (hs *HTTPServer) addUserRoutes() {
	r := hs.RouteRegister

	r.Post("/signin", hs.signin)
	r.Post("/signup", hs.signup)
}

func (hs *HTTPServer) signin(ctx web.Context) error {
	fmt.Println(ctx.Path())
	return ctx.String(http.StatusOK, "Hello login\n")
}

func (hs *HTTPServer) signup(ctx web.Context) error {
	c := &command.RegisterUserCommand{}

	if err := ctx.Bind(c); err != nil {
		return err
	}

	err := hs.UserApp.Commands.RegiseterUser.Handle(*c)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, c)
}
