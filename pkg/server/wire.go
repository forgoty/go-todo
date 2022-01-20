// +build wireinject

package server

import (
	userapp "github.com/forgoty/go-todo/internal/user/app"
	"github.com/forgoty/go-todo/pkg/api"
	"github.com/google/wire"
)

func ProvideServer(port string) (*Server, error) {
	wire.Build(NewProvide, api.ProvideHTTPServer, userapp.NewUserApplication)
	return &Server{}, nil
}
