package api

import (
	"net/http"

	"github.com/forgoty/go-todo/internal/user/app/command"
	"github.com/forgoty/go-todo/internal/user/app/query"
	"github.com/forgoty/go-todo/pkg/models"
	"github.com/forgoty/go-todo/pkg/web"
	"github.com/google/uuid"
)

func (hs *HTTPServer) addUserRoutes() {
	r := hs.RouteRegister

	r.Post("/signin", hs.signin)
	r.Post("/signup", hs.signup)
}

func (hs *HTTPServer) signin(ctx web.Context) error {
	dto := &models.UserSignInSignUp{}
	if err := ctx.Bind(dto); err != nil {
		return err
	}
	encryptedPassword := hs.AuthService.GeneratePasswordHash(dto.Password)

	q := query.FindUserBySigninQuery{
		Username: dto.Username,
		Password: encryptedPassword,
	}
	u, err := hs.UserApp.Queries.FindUserBySignin.Handle(q)
	if err != nil {
		return err
	}
	token, err := hs.AuthService.GenerateToken(u.Id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}

func (hs *HTTPServer) signup(ctx web.Context) error {
	dto := &models.UserSignInSignUp{}

	if err := ctx.Bind(dto); err != nil {
		return err
	}

	encryptedPassword := hs.AuthService.GeneratePasswordHash(dto.Password)

	id := uuid.New().String()
	c := &command.RegisterUserCommand{
		Id:           id,
		Username:     dto.Username,
		PasswordHash: encryptedPassword,
	}

	err := hs.UserApp.Commands.RegiseterUser.Handle(*c)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusCreated)
}
