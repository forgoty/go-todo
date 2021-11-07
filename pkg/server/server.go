package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/wire"
	"sync"

	"github.com/forgoty/go-todo/pkg/api"
)

type Server struct {
	port             string
	context          context.Context
	shutdownFn       context.CancelFunc
	shutdownOnce     sync.Once
	shutdownFinished chan struct{}
	isInitialized    bool
	mtx              sync.Mutex
	version          string
	HTTPServer       *api.HTTPServer
}

var wireBasicSet = wire.NewSet(
	api.ProvideHTTPServer,
)

func Initialize(port string) (*Server, error) {
	wire.Build(wireBasicSet)
	return &Server{}, nil
}

func New(port string) (*Server, error) {
	httpServer, err := api.ProvideHTTPServer()
	if err != nil {
		return &Server{}, err
	}
	rootCtx, shutdownFn := context.WithCancel(context.Background())

	s := &Server{
		port:             port,
		context:          rootCtx,
		HTTPServer:       httpServer,
		shutdownFn:       shutdownFn,
		shutdownFinished: make(chan struct{}),
	}
	s.isInitialized = true

	return s, nil
}

func (s *Server) Run() error {
	defer close(s.shutdownFinished)

	err := s.HTTPServer.Run(s.context, s.port)
	if err != nil && !errors.Is(err, context.Canceled) {
		fmt.Printf("Stopped background service. Reason %s\n", err.Error())
		return fmt.Errorf("run error: %w\n", err)
	}
	return err
}

func (s *Server) Shutdown(ctx context.Context, reason string) error {
	var err error
	s.shutdownOnce.Do(func() {
		fmt.Printf("\nShutdown started. Reason: %s\n", reason)
		// Call cancel func to stop services.
		s.shutdownFn()
		// Wait for server to shut down
		select {
		case <-s.shutdownFinished:
			fmt.Println("Finished waiting for server to shut down")
		case <-ctx.Done():
			fmt.Println("Timed out while waiting for server to shut down")
			err = fmt.Errorf("timeout waiting for shutdown\n")
		}
	})

	return err
}
