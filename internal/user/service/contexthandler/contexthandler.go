package contexthandler

import (
	"errors"
	"strings"

	userapp "github.com/forgoty/go-todo/internal/user/app"
	userquery "github.com/forgoty/go-todo/internal/user/app/query"
	"github.com/forgoty/go-todo/internal/user/service/auth"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
	"github.com/forgoty/go-todo/pkg/web"
)

type ContextHandler struct {
	userApp    *userapp.Application // replace with a Bus
	jwtService *auth.JWTService
	logger     logger.Logger
}

type SignedInUser struct {
	UserId      string
	Username    string
	IsAnonymous bool
}

const UserKey = "user"

func NewContextHandler(u *userapp.Application, jwtService *auth.JWTService) *ContextHandler {
	return &ContextHandler{
		userApp:    u,
		logger:     logger.New("contexthandler"),
		jwtService: jwtService,
	}
}

func (h *ContextHandler) Middleware(next web.HandlerFunc) web.HandlerFunc {
	return func(c web.Context) error {
		switch {
		case h.initContextWithJWT(c):
		case h.initContextWithAnonymousUser(c):
		}
		return next(c)
	}
}

func (h *ContextHandler) initContextWithJWT(ctx web.Context) bool {
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

	userId, err := h.jwtService.ParseToken(headerParts[1])
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

	ctx.Set(UserKey, SignedInUser{
		UserId:      user.Id,
		Username:    user.Username,
		IsAnonymous: false,
	})

	return true
}

func (h *ContextHandler) initContextWithAnonymousUser(ctx web.Context) bool {
	ctx.Set(UserKey, SignedInUser{
		IsAnonymous: true,
	})
	return true
}
