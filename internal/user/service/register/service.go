package register

import (
	"errors"

	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/internal/user/domain/user/commands"
	"github.com/forgoty/go-todo/internal/user/service/auth"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type RegisterService struct {
	log         logger.Logger
	userRepo    aggregates.IUserRepository
	authService *auth.AuthService
}

func NewRegisterService(userRepo aggregates.IUserRepository, as *auth.AuthService) *RegisterService {
	log := logger.New("registerService")
	return &RegisterService{
		log:         log,
		userRepo:    userRepo,
		authService: as,
	}
}

func (rs *RegisterService) RegisterUser(c commands.RegisterUserCommand) error {
	encryptedPassword := rs.authService.GeneratePasswordHash(c.Password)
	user, err := aggregates.NewUser(c, encryptedPassword)
	if err != nil {
		return err
	}
	if err = rs.userRepo.Create(*user); err != nil {
		return errors.New("Cannot Create user")
	}
	return nil
}
