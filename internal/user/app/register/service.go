package register

import (
	"errors"

	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/internal/user/service/auth"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type RegisterService struct {
	log      logger.Logger
	userRepo aggregates.IUserRepository
	passMgr  *auth.PasswordManager
}

func NewRegisterService(userRepo aggregates.IUserRepository, pswmgr *auth.PasswordManager) *RegisterService {
	log := logger.New("registerService")
	return &RegisterService{
		log:      log,
		userRepo: userRepo,
		passMgr:  pswmgr,
	}
}

func (rs *RegisterService) RegisterUser(c RegisterUserCommand) error {
	encryptedPassword := rs.passMgr.HashPassword(c.Password)
	user, err := aggregates.NewUser(c.Id, c.Username, c.Password, encryptedPassword)
	if err != nil {
		return err
	}
	if err = rs.userRepo.Create(*user); err != nil {
		return errors.New("Cannot Create user")
	}
	return nil
}
