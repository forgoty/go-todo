package query

import (
	"github.com/forgoty/go-todo/internal/user/app/query/models"
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

const (
	ANON = iota
	SIGNED
	SELF
)

type GetUserQuery struct {
	Id   string
	Mode uint
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

// TODO Make hiden fields calculated in api layer
func (h *GetUserQueryHandler) Handle(q GetUserQuery) (*models.UserDto, error) {
	u, err := h.userRepo.FindOneById(q.Id)
	if err != nil {
		return nil, err
	}
	dto := &models.UserDto{
		Id:        u.Id,
		Username:  string(u.Username),
		FirstName: u.UserProfile.FirstName,
		LastName:  u.UserProfile.LastName,
	}
	switch q.Mode {
	case SIGNED:
		{
			dto.Personal = u.UserProfile.PersonalField
		}
	case SELF:
		{
			dto.Personal = u.UserProfile.PersonalField
			dto.Secret = u.UserProfile.SecretField
		}
	}
	return dto, nil
}
