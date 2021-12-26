package api

import (
	"net/http"

	"github.com/forgoty/go-todo/internal/user/app/command"
	"github.com/forgoty/go-todo/internal/user/app/query"
	"github.com/forgoty/go-todo/pkg/web"
	"github.com/google/uuid"
)

func (hs *HTTPServer) addUserRoutes() {
	r := hs.RouteRegister

	r.Post("/signin", hs.signin)
	r.Post("/signup", hs.signup)
}

func (hs *HTTPServer) signin(ctx web.Context) error {
	q := &query.GetUserQuery{}
	if err := ctx.Bind(q); err != nil {
		return err
	}
	u, err := hs.UserApp.Queries.GetUser.Handle(*q)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, u)
}

func (hs *HTTPServer) signup(ctx web.Context) error {
	id := uuid.New().String()
	c := &command.RegisterUserCommand{
		Id: id,
	}

	if err := ctx.Bind(c); err != nil {
		return err
	}

	err := hs.UserApp.Commands.RegiseterUser.Handle(*c)
	if err != nil {
		return err
	}
	return ctx.String(http.StatusCreated, id)
}
