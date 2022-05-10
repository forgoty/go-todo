package query

import (
	"github.com/forgoty/go-todo/internal/user/app/query/models"
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type GetUserQuery struct {
	Id string `json:"id"`
}

type GetUserQueryHandler struct {
	log      logger.Logger
	userRepo aggregates.IUserRepository
}

func ProvideGetUserQueryHandler(repo aggregates.IUserRepository) *GetUserQueryHandler {
	return &GetUserQueryHandler{
		log:      logger.New("GetUserQueryHandler "),
		userRepo: repo,
	}
}

func (h *GetUserQueryHandler) Handle(q GetUserQuery) (*models.UserDto, error) {
	u, err := h.userRepo.FindOneById(q.Id)
	if err != nil {
		return nil, err
	}
	return &models.UserDto{
		Username: string(u.Username),
	}, nil
}
