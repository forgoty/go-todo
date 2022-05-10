package rest

import (
	"net/http"

	"github.com/forgoty/go-todo/internal/user/domain/user/commands"
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
		return ctx.JSON(http.StatusBadRequest, api_models.APIError{Message: err.Error()})
	}

	command := commands.LoginUserWithJWTCommand{
		Username: dto.Username,
		Password: dto.Password,
	}
	token, err := c.userApp.Commands.LoginUserJWT.Handle(command)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, api_models.APIError{Message: err.Error()})
	}
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
		return ctx.JSON(http.StatusBadRequest, api_models.APIError{Message: err.Error()})
	}

	command := commands.RegisterUserCommand{
		Id:       uuid.New().String(),
		Username: dto.Username,
		Password: dto.Password,
	}

	err := c.userApp.Commands.RegiseterUser.Handle(command)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, api_models.APIError{Message: err.Error()})
	}
	return ctx.NoContent(http.StatusCreated)
}
