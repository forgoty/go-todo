package server

import (
	"context"
	"errors"
	"fmt"
	"sync"

	userapp "github.com/forgoty/go-todo/internal/user/app"
	"github.com/forgoty/go-todo/pkg/api"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
)

type Server struct {
	port             string
	context          context.Context
	shutdownFn       context.CancelFunc
	shutdownOnce     sync.Once
	shutdownFinished chan struct{}
	isInitialized    bool
	HTTPServer       *api.HTTPServer
	log              logger.Logger
}

func New(port string) (*Server, error) {
	// We could add some new backgraund services here in future
	userApp := userapp.NewUserApplication()
	httpServer, err := api.ProvideHTTPServer(userApp)
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
		log:              logger.New("server"),
	}
	s.isInitialized = true

	return s, nil
}

func (s *Server) Run() error {
	defer close(s.shutdownFinished)

	err := s.HTTPServer.Run(s.context, s.port)
	if err != nil && !errors.Is(err, context.Canceled) {
		s.log.Error("Stopped background service. Reason %s", err.Error())
		return fmt.Errorf("run error: %w", err)
	}
	return err
}

func (s *Server) Shutdown(ctx context.Context, reason string) error {
	var err error
	s.shutdownOnce.Do(func() {
		s.log.Info("Shutdown started. Reason: %s", reason)
		// Call cancel func to stop services.
		s.shutdownFn()
		// Wait for server to shut down
		select {
		case <-s.shutdownFinished:
			s.log.Info("Finished waiting for server to shut down")
		case <-ctx.Done():
			s.log.Error("Timed out while waiting for server to shut down")
			err = fmt.Errorf("timeout waiting for shutdown")
		}
	})

	return err
}
