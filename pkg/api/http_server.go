package api

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/forgoty/go-todo/pkg/api/routing"
	"github.com/forgoty/go-todo/pkg/web"
)

type HTTPServer struct {
	context       context.Context
	httpSrv       *http.Server
	RouteRegister routing.RouteRegister
	web           *web.Handler
}

func ProvideHTTPServer() (*HTTPServer, error) {
	return &HTTPServer{
		httpSrv:       nil,
		web:           web.New(),
		RouteRegister: &routing.RouteRegisterImpl{},
	}, nil
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
			fmt.Printf("Failed to shutdown server. Error: %s\n", err.Error())
		}
	}()

	if err := hs.httpSrv.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server was shutdown gracefully")
			return nil
		}
		return err
	}

	wg.Wait()

	return nil
}

func (hs *HTTPServer) applyRoutes() {
	hs.RouteRegister.Register(hs.web.Router())
	handle := func(ctx web.Context) error {
		return ctx.String(http.StatusOK, "Hello world\n")
	}

	hs.RouteRegister.Get("/", handle)
}
