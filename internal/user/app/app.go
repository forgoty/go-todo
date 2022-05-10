package app

import (
	"github.com/forgoty/go-todo/internal/user/app/command"
	"github.com/forgoty/go-todo/internal/user/app/query"
)

type Application struct {
	Commands *Commands
	Queries  *Queries
}

type Commands struct {
	RegiseterUser *command.RegisterUserCommandHandler
	LoginUserJWT  *command.LoginUserWithJWTCommandHandler
}

type Queries struct {
	FindUserBySignin *query.FindUserBySigninQueryHandler
	GetUser          *query.GetUserQueryHandler
}
