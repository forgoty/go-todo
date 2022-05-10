package query

import (
	"github.com/forgoty/go-todo/internal/user/app/query/models"
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type FindUserBySigninQuery struct {
	Username string
	Password string
}

type FindUserBySigninQueryHandler struct {
	log      logger.Logger
	userRepo aggregates.IUserRepository
}

func ProvideFindUserBySigninQueryHandler(repo aggregates.IUserRepository) *FindUserBySigninQueryHandler {
	return &FindUserBySigninQueryHandler{
		log:      logger.New("FindUserBySigninQueryHandler "),
		userRepo: repo,
	}
}

func (h FindUserBySigninQueryHandler) Handle(q FindUserBySigninQuery) (*models.UserDto, error) {
	u, err := h.userRepo.FindOneByUsernameAndPassword(q.Username, q.Password)
	if err != nil {
		return nil, err
	}
	return &models.UserDto{
		Id:       u.Id,
		Username: string(u.Username),
	}, nil
}
