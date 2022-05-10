#!/bin/sh
makewire() {
	wire pkg/server/*.go &&
	wire internal/user/interfaces/rest/*.go
	wire internal/user/app/*.go
}


makewire
