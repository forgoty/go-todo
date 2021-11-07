package commands

import (
	"flag"
	"fmt"
	"os"
)

var serverFlagSet = flag.NewFlagSet("todo-server", flag.ContinueOnError)

func RunCli(versionArg string) int {
	var (
		port    = serverFlagSet.Int("port", 8000, "Provide server port")
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
	fmt.Printf("Server runs on port %d\n", *port)

	return 0
}
