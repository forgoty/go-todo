package query

import (
	"github.com/forgoty/go-todo/internal/user/domain/user"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type GetUserQuery struct {
	Id string `json:"id"`
}

type UserDto struct {
	Username string `json:"username"`
}

type GetUserQueryHandler struct {
	log      logger.Logger
	userRepo user.IUserRepository
}

func NewGetUserQueryHandler(repo user.IUserRepository, log logger.Logger) *GetUserQueryHandler {
	return &GetUserQueryHandler{
		userRepo: repo,
		log:      log,
	}
}

func (h *GetUserQueryHandler) Handle(q GetUserQuery) (*UserDto, error) {
	u, err := h.userRepo.FindOneById(q.Id)
	if err != nil {
		return nil, err
	}
	return &UserDto{
		Username: u.Username,
	}, nil
}
