package login

import (
	"errors"

	"github.com/forgoty/go-todo/internal/user/domain/user/aggregates"
	"github.com/forgoty/go-todo/internal/user/domain/user/commands"
	"github.com/forgoty/go-todo/internal/user/service/auth"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

var ErrInvalidPasswordHash = errors.New("No account found with the given credentials")

type LoginService struct {
	log        logger.Logger
	userRepo   aggregates.IUserRepository
	passMgr    *auth.PasswordManager
	jwtService *auth.JWTService
}

func NewLoginService(userRepo aggregates.IUserRepository, passMgr *auth.PasswordManager, jwt *auth.JWTService) *LoginService {
	log := logger.New("loginService")
	return &LoginService{
		log:        log,
		userRepo:   userRepo,
		passMgr:    passMgr,
		jwtService: jwt,
	}
}

func (ls *LoginService) LoginUserJWT(c commands.LoginUserWithJWTCommand) (string, error) {
	user, err := ls.userRepo.FindOneByUsername(c.Username)
	if err != nil {
		return "", err
	}
	if !ls.passMgr.VerifyHashedPassword(c.Password, user.PasswordHash) {
		return "", ErrInvalidPasswordHash
	}
	token, err := ls.jwtService.GenerateToken(user.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}
