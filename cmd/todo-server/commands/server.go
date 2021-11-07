package commands

import (
	"context"
	"sync"
	// "github.com/google/wire"
	// "github.com/forgoty/go-todo/pkg/api"
)

type Server struct {
	context          context.Context
	shutdownFn       context.CancelFunc
	shutdownOnce     sync.Once
	shutdownFinished chan struct{}
	isInitialized    bool
	mtx              sync.Mutex
	version          string
	//HTTPServer       *api.HTTPServer
}

func Initialize(port int) (*Server, error) {
	//wire.Build(wireSet)
	return &Server{}, nil
}
