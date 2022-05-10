package rest

import (
	"time"

	"github.com/forgoty/go-todo/internal/user/app"
	"github.com/forgoty/go-todo/internal/user/service/contexthandler"
	"github.com/forgoty/go-todo/pkg/api/routing"
	"github.com/forgoty/go-todo/pkg/web"
)

type userController struct {
	userApp       *app.Application
	contextHander *contexthandler.ContextHandler
}

func RegisterRoutesAndMiddlewares(r routing.RouteRegister, m *web.Handler) {
	c, _ := provideUserController("123", "123", 12*time.Hour)

	m.Use(c.contextHander.Middleware)

	r.Post("/signin", c.signin)
	r.Post("/signup", c.signup)
}
