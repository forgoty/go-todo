package api

import (
	"context"
	"errors"
	"net"
	"net/http"
	"sync"
	"time"

	userapp "github.com/forgoty/go-todo/internal/user/app"
	"github.com/forgoty/go-todo/pkg/api/routing"
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
	"github.com/forgoty/go-todo/pkg/services/auth"
	"github.com/forgoty/go-todo/pkg/services/contexthandler"
	"github.com/forgoty/go-todo/pkg/web"
)

type HTTPServer struct {
	context        context.Context
	httpSrv        *http.Server
	RouteRegister  routing.RouteRegister
	ContextHandler *contexthandler.ContextHandler
	AuthService    *auth.AuthService
	web            *web.Handler
	log            logger.Logger
	UserApp        *userapp.Application
}

func ProvideHTTPServer(userApp *userapp.Application) (*HTTPServer, error) {
	authService := auth.NewAuthService("123", "123", 12*time.Hour)
	contextHandler := contexthandler.NewContextHandler(userApp, authService, logger.New("contexthandler"))
	hs := &HTTPServer{
		httpSrv:        nil,
		web:            web.New(),
		RouteRegister:  routing.NewRouteRegister(),
		ContextHandler: contextHandler,
		AuthService:    authService,
		log:            logger.New("httpserver"),
		UserApp:        userApp,
	}
	hs.registerRoutes()

	return hs, nil
}

func (hs *HTTPServer) Run(ctx context.Context, port string) error {
	hs.context = ctx

	hs.applyRoutes()
	hs.httpSrv = &http.Server{
		Addr:    net.JoinHostPort("localhost", port),
		Handler: hs.web,
	}
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		<-ctx.Done()
		if err := hs.httpSrv.Shutdown(context.Background()); err != nil {
			hs.log.Error("Failed to shutdown server. Error: %s\n", err.Error())
		}
	}()

	if err := hs.httpSrv.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			hs.log.Info("Server was shutdown gracefully")
			return nil
		}
		return err
	}

	wg.Wait()

	return nil
}

func (hs *HTTPServer) applyRoutes() {
	hs.addMiddlewares()
	hs.RouteRegister.Register(hs.web.Router())
}

func (hs *HTTPServer) addMiddlewares() {
	m := hs.web
	m.Use(hs.ContextHandler.Middleware)
	m.Use(web.MiddlewareLogger())
	m.Use(web.MiddlewareRecover())
}
