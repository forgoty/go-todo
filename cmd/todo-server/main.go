package main

import (
	"os"

	"github.com/forgoty/go-todo/cmd/todo-server/commands"
)

var version = "0.0.1"

// @Version            1.0.0
// @Title              ToDo application
// @Description        Basic ToDo Application.
// @ContactName        forgoty
// @ContactEmail       example.com
// @ContactURL         http://example.com
// @TermsOfServiceUrl  http://example.com
// @LicenseName        MIT
// @LicenseURL         https://en.wikipedia.org/wiki/MIT_License
// @Server             http://www.example.com Server-1
// @Server             http://www.example.com Server-2
// @Security           AuthorizationHeader read write
// @SecurityScheme     AuthorizationHeader bearer
func main() {
	os.Exit(commands.RunCli(version))
}
