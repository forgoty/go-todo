package login

import (
	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/internal/user/domain/user/commands"
	"github.com/forgoty/go-todo/internal/user/service/auth"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type LoginService struct {
	log         logger.Logger
	userRepo    aggregates.IUserRepository
	authService *auth.AuthService
}

func NewLoginService(userRepo aggregates.IUserRepository, as *auth.AuthService) *LoginService {
	log := logger.New("loginService")
	return &LoginService{
		log:         log,
		userRepo:    userRepo,
		authService: as,
	}
}

func (ls *LoginService) LoginUserJWT(c commands.LoginUserWithJWTCommand) (string, error) {
	pHash := ls.authService.GeneratePasswordHash(c.Password)
	user, err := ls.userRepo.FindOneByUsernameAndPassword(c.Username, pHash)
	if err != nil {
		return "", err
	}
	token, err := ls.authService.GenerateToken(user.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}
