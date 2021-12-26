package app

import (
	"github.com/forgoty/go-todo/internal/user/app/command/handler"
	"github.com/forgoty/go-todo/internal/user/infrastructure/persistance"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type Application struct {
	Commands *Commands
	Queries  *Queries
}

type Commands struct {
	RegiseterUser *handler.RegisterUserHandler
}

type Queries struct {
}

func NewUserApplication() *Application {
	user_repo := persistance.NewInMemoryUserRepository()
	commands := &Commands{
		RegiseterUser: handler.NewRegisterUserHandler(
			user_repo,
			logger.New("registeruser"),
		),
	}
	queries := &Queries{}
	return &Application{
		Commands: commands,
		Queries:  queries,
	}
}
