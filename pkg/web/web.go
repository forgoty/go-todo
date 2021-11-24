package web

import (
	"github.com/labstack/echo/v4"
)

type Context = echo.Context
type Handler = echo.Echo
type Mux = echo.Router
type HandlerFunc = echo.HandlerFunc
type MiddlewareFunc = echo.MiddlewareFunc

var New = echo.New
