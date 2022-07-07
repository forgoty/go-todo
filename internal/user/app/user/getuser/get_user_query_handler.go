package getuser

import (
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type GetUserQueryHandler struct {
	log      logger.Logger
	userRepo aggregates.IUserRepository
}

func (h *GetUserQueryHandler) Handle(q GetUserQuery) (*UserDto, error) {
	u, err := h.userRepo.FindOneById(q.Id)
	if err != nil {
		return nil, err
	}
	dto := &UserDto{
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
