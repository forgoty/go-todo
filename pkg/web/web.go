package web

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Context = echo.Context
type Handler = echo.Echo
type Mux = echo.Router
type HandlerFunc = echo.HandlerFunc
type MiddlewareFunc = echo.MiddlewareFunc
type Response = echo.Response
type Logger = echo.Logger

var New = echo.New
var SwaggerHander = echoSwagger.WrapHandler
