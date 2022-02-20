package rest

import (
	"net/http"
	"time"

	"github.com/forgoty/go-todo/internal/user/app"
	"github.com/forgoty/go-todo/internal/user/app/command"
	"github.com/forgoty/go-todo/internal/user/app/query"
	"github.com/forgoty/go-todo/internal/user/interfaces/rest/models"
	"github.com/forgoty/go-todo/internal/user/service/auth"
	"github.com/forgoty/go-todo/internal/user/service/contexthandler"
	"github.com/forgoty/go-todo/pkg/api/routing"
	"github.com/forgoty/go-todo/pkg/web"
	"github.com/google/uuid"
)

func RegisterRoutesAndMiddlewares(r routing.RouteRegister, m *web.Handler) {
	c, _ := provideUserController("123", "123", 12*time.Hour)

	m.Use(c.contextHander.Middleware)

	r.Post("/signin", c.signin)
	r.Post("/signup", c.signup)
}

type userController struct {
	authService   *auth.AuthService
	userApp       *app.Application
	contextHander *contexthandler.ContextHandler
}

func (c *userController) signin(ctx web.Context) error {
	dto := &models.UserSignInSignUp{}
	if err := ctx.Bind(dto); err != nil {
		return err
	}
	encryptedPassword := c.authService.GeneratePasswordHash(dto.Password)

	q := query.FindUserBySigninQuery{
		Username: dto.Username,
		Password: encryptedPassword,
	}
	u, err := c.userApp.Queries.FindUserBySignin.Handle(q)
	if err != nil {
		return err
	}
	token, err := c.authService.GenerateToken(u.Id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}

func (c *userController) signup(ctx web.Context) error {
	dto := &models.UserSignInSignUp{}

	if err := ctx.Bind(dto); err != nil {
		return err
	}

	encryptedPassword := c.authService.GeneratePasswordHash(dto.Password)

	id := uuid.New().String()
	command := &command.RegisterUserCommand{
		Id:           id,
		Username:     dto.Username,
		PasswordHash: encryptedPassword,
	}

	err := c.userApp.Commands.RegiseterUser.Handle(*command)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusCreated)
}
