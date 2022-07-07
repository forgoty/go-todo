package getuser

import (
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

func ProvideGetUserQueryHandler(repo aggregates.IUserRepository) *GetUserQueryHandler {
	return &GetUserQueryHandler{
		log:      logger.New("GetUserQueryHandler "),
		userRepo: repo,
	}
}
