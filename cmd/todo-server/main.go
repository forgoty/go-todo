package main

import (
	"os"

	"github.com/forgoty/go-todo/cmd/todo-server/commands"
)

var version = "0.0.1"

func main() {
	os.Exit(commands.RunCli(version))
}
