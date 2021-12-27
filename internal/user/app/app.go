package app

import (
	"github.com/forgoty/go-todo/internal/user/app/command"
	"github.com/forgoty/go-todo/internal/user/app/query"
	"github.com/forgoty/go-todo/internal/user/infrastructure/persistance"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type Application struct {
	Commands *Commands
	Queries  *Queries
}

type Commands struct {
	RegiseterUser *command.RegisterUserCommandHandler
}

type Queries struct {
	FindUserBySignin *query.FindUserBySigninQueryHandler
	GetUser          *query.GetUserQueryHandler
}

func NewUserApplication() *Application {
	user_repo := persistance.NewInMemoryUserRepository()
	commands := &Commands{
		RegiseterUser: command.NewRegisterUserCommandHandler(
			user_repo,
			logger.New("registeruser"),
		),
	}
	queries := &Queries{
		FindUserBySignin: query.NewFindUserBySigninQueryHandler(user_repo, logger.New("finduserbysignin")),
		GetUser:          query.NewGetUserQueryHandler(user_repo, logger.New("getuser")),
	}
	return &Application{
		Commands: commands,
		Queries:  queries,
	}
}
