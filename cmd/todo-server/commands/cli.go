package commands

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/forgoty/go-todo/pkg/server"
)

var serverFlagSet = flag.NewFlagSet("todo-server", flag.ContinueOnError)

func RunCli(versionArg string) int {
	var (
		port    = serverFlagSet.String("port", "8000", "Provide server port")
		version = serverFlagSet.Bool("version", false, "print current version and exists")
	)

	if err := serverFlagSet.Parse(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	if *version {
		fmt.Printf("Version is %s\n", versionArg)
		return 0
	}
	s, err := server.ProvideServer(*port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start todo server. error: %s\n", err)
		return 1
	}

	ctx := context.Background()

	go listenToSystemSignals(ctx, s)

	if err := s.Run(); err != nil {
		return 1
	}

	return 0
}

func listenToSystemSignals(ctx context.Context, s *server.Server) {
	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case sig := <-signalChan:
			ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()
			if err := s.Shutdown(ctx, fmt.Sprintf("System signal: %s", sig)); err != nil {
				fmt.Fprintf(os.Stderr, "Timed out waiting for server to shut down\n")
			}
			return
		}
	}
}
