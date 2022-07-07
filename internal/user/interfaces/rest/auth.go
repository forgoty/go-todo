package rest

import (
	"errors"
	"net/http"

	"github.com/forgoty/go-todo/internal/user/app/login"
	"github.com/forgoty/go-todo/internal/user/app/register"
	"github.com/forgoty/go-todo/internal/user/interfaces/rest/models"
	api_models "github.com/forgoty/go-todo/pkg/api/models"
	"github.com/forgoty/go-todo/pkg/errorcollector"
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
	collector := make(errorcollector.Fields)
	if dto.Username == "" {
		collector.Add("username", errors.New("field not provided"))
	}
	if dto.Password == "" {
		collector.Add("password", errors.New("field not provided"))
	}
	if err := collector.Err(); err != nil {
		return ctx.JSON(http.StatusBadRequest, api_models.APIError{Message: err.Error()})
	}

	command := login.LoginUserWithJWTCommand{
		Username: dto.Username,
		Password: dto.Password,
	}
	token, err := c.userApp.Commands.LoginUserJWT.Handle(command)
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
	collector := make(errorcollector.Fields)
	if dto.Username == "" {
		collector.Add("username", errors.New("field not provided"))
	}
	if dto.Password == "" {
		collector.Add("password", errors.New("field not provided"))
	}
	if err := collector.Err(); err != nil {
		return ctx.JSON(http.StatusBadRequest, api_models.APIError{Message: err.Error()})
	}

	command := register.RegisterUserCommand{
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
