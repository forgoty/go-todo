package rest

import (
	"net/http"

	"github.com/forgoty/go-todo/internal/user/app/command"
	"github.com/forgoty/go-todo/internal/user/app/query"
	"github.com/forgoty/go-todo/internal/user/interfaces/rest/models"
	api_models "github.com/forgoty/go-todo/pkg/api/models"
	"github.com/forgoty/go-todo/pkg/web"
	"github.com/google/uuid"
)

// @Summary      User sign-in
// @Description  User sign-in
// @Accept       json
// @Produce      json
// @Param        UserSignIn  body      models.UserSignInSignUp  true  "Username and password"
// @Success      200         {object}  models.Token             "JWT Token"
// @Failure      400         {object}  api_models.APIError      "API Error"
// @BasePath     api/v1
// @Router       /api/v1/signin [post]
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
		return ctx.JSON(http.StatusBadRequest, api_models.APIError{Message: err.Error()})
	}
	token, err := c.authService.GenerateToken(u.Id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, api_models.APIError{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Token{Token: token})
}

// @Summary      User sign-up
// @Description  User sign-up by username and password.
// @Accept       json
// @Produce      json
// @Param        UserSignUp  body  models.UserSignInSignUp  true  "Username and password"
// @Success      201         "User signed up successfully"
// @Failure      400         {object}  api_models.APIError  "API Error"
// @BasePath     api/v1
// @Router       /api/v1/signup [post]
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
		return ctx.JSON(http.StatusBadRequest, api_models.APIError{Message: err.Error()})
	}
	return ctx.NoContent(http.StatusCreated)
}
