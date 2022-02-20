// +build wireinject

package server

import (
	"github.com/forgoty/go-todo/pkg/api"
	"github.com/google/wire"
)

func ProvideServer(port string) (*Server, error) {
	wire.Build(NewServer, api.ProvideHTTPServer)
	return &Server{}, nil
}
