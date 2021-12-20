package app

import (
	"github.com/forgoty/go-todo/internal/user/app/command/handler"
)

type Application struct {
	Commands *Commands
	Queries  *Queries
}

type Commands struct {
	RegiseterUser handler.RegisterUserHandler
}

type Queries struct {
}

func NewUserApplication() *Application {
	commands := &Commands{
		RegiseterUser: handler.RegisterUserHandler{},
	}
	queries := &Queries{}
	return &Application{
		Commands: commands,
		Queries:  queries,
	}
}
