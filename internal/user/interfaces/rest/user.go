package rest

import (
	"net/http"

	"github.com/forgoty/go-todo/internal/user/app/query"
	"github.com/forgoty/go-todo/internal/user/service/contexthandler"
	api_models "github.com/forgoty/go-todo/pkg/api/models"
	"github.com/forgoty/go-todo/pkg/web"
)

// @Summary      Get user by Id
// @Description  Get user by Id.
// @Accept       json
// @Produce      json
// @Success      200         {object}  query_models.UserDto "User found"
// @BasePath     api/v1/user
// @Router       /api/v1/user/{id} [get]
func (c *userController) getUserById(ctx web.Context) error {
	if ctx.Param("id") == "" {
		return ctx.JSON(http.StatusBadRequest, api_models.APIError{Message: "Could not parse field: id"})
	}
	requestUser := ctx.Get(contexthandler.UserKey).(contexthandler.SignedInUser)
	var mode uint
	switch {
	case requestUser.IsAnonymous:
		mode = query.ANON
	case requestUser.UserId == ctx.Param("id"):
		mode = query.SELF
	default:
		mode = query.SIGNED
	}

	userQuery := query.GetUserQuery{
		Id:   ctx.Param("id"),
		Mode: mode,
	}
	user, err := c.userApp.Queries.GetUser.Handle(userQuery)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, api_models.APIError{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, user)
}
