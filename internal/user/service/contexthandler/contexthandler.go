package contexthandler

import (
	"errors"
	"strings"

	userapp "github.com/forgoty/go-todo/internal/user/app"
	userquery "github.com/forgoty/go-todo/internal/user/app/query"
	"github.com/forgoty/go-todo/internal/user/service/auth"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
	"github.com/forgoty/go-todo/pkg/models"
	"github.com/forgoty/go-todo/pkg/web"
)

type ContextHandler struct {
	userApp     *userapp.Application // replace with a Bus
	authService *auth.AuthService
	logger      logger.Logger
}

func NewContextHandler(u *userapp.Application, a *auth.AuthService) *ContextHandler {
	return &ContextHandler{
		userApp:     u,
		authService: a,
		logger:      logger.New("contexthandler"),
	}
}

func (h *ContextHandler) Middleware(next web.HandlerFunc) web.HandlerFunc {
	return func(c web.Context) error {
		reqContext := &models.ReqContext{
			Context: c,
		}

		switch {
		case h.initContextWithJWT(reqContext):
		case h.initContextWithAnonymousUser(reqContext):
		}
		return next(reqContext)
	}
}

func (h *ContextHandler) initContextWithJWT(ctx *models.ReqContext) bool {
	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		return false
	}
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		e := errors.New("No Bearer part or no token")
		h.logger.Debug("Failed to parse JWT", "error", e)
		if resErr := ctx.JSON(401, "Unauthorized"); resErr != nil {
			h.logger.Debug("Failed to send response", "error", resErr)
		}
	}

	userId, err := h.authService.ParseToken(headerParts[1])
	if err != nil {
		h.logger.Debug("Failed to verify JWT", "error", err)
		if resErr := ctx.JSON(401, "Unauthorized"); resErr != nil {
			h.logger.Debug("Failed to send response", "error", resErr)
		}
		return true
	}

	// TODO: Send via Bus
	query := userquery.GetUserQuery{Id: userId}
	user, err := h.userApp.Queries.GetUser.Handle(query)
	if err != nil {
		h.logger.Debug("Failed to find user using JWT claims", "error", err)
		if resErr := ctx.JSON(401, "Unauthorized"); resErr != nil {
			h.logger.Debug("Failed to send response", "error", resErr)
		}
		return true
	}

	ctx.SignedInUser = &models.SignedInUser{
		UserId:      user.Id,
		Username:    user.Username,
		IsAnonymous: false,
	}
	ctx.IsSignedIn = true

	return true
}

func (h *ContextHandler) initContextWithAnonymousUser(reqContext *models.ReqContext) bool {
	reqContext.IsSignedIn = false
	reqContext.AllowAnonymous = true
	reqContext.SignedInUser = &models.SignedInUser{IsAnonymous: true}
	return true
}
