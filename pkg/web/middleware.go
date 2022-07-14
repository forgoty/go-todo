package web

import (
	"github.com/labstack/echo/v4/middleware"
)

var MiddlewareLogger = middleware.Logger
var MiddlewareRecover = middleware.Recover
var MiddlewareCORS = middleware.CORS
