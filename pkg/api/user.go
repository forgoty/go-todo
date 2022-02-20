package api

import "github.com/forgoty/go-todo/internal/user/interfaces/rest"

func (hs *HTTPServer) addUserRoutes() {
	rest.RegisterRoutesAndMiddlewares(hs.routeRegister, hs.web)
}
