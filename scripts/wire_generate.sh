#!/bin/sh
makewire() {
	wire pkg/server/*.go &&
	wire internal/user/interfaces/rest/*.go
}


makewire 2> /dev/null || echo "Files already generated"
