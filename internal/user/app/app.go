package app

import (
	"github.com/forgoty/go-todo/internal/user/app/login"
	"github.com/forgoty/go-todo/internal/user/app/register"
	"github.com/forgoty/go-todo/internal/user/app/user/getuser"
)

type Application struct {
	Commands *Commands
	Queries  *Queries
}

type Commands struct {
	RegiseterUser *register.RegisterUserCommandHandler
	LoginUserJWT  *login.LoginUserWithJWTCommandHandler
}

type Queries struct {
	GetUser *getuser.GetUserQueryHandler
}
