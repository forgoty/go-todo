package routing

import (
	"net/http"

	"github.com/forgoty/go-todo/pkg/web"
)

// Echo provide Add method to register the routes
type Router interface {
	Add(method, pattern string, handler web.HandlerFunc)
}

// RouteRegister allows you to add routes and web.Handlers
// that the web server should serve.
type RouteRegister interface {
	// Get adds a list of handler to a given route with a GET HTTP verb
	Get(string, web.HandlerFunc)

	// Post adds a list of handler to a given route with a POST HTTP verb
	Post(string, web.HandlerFunc)

	// Delete adds a list of handler to a given route with a DELETE HTTP verb
	Delete(string, web.HandlerFunc)

	// Put adds a list of handler to a given route with a PUT HTTP verb
	Put(string, web.HandlerFunc)

	// Patch adds a list of handler to a given route with a PATCH HTTP verb
	Patch(string, web.HandlerFunc)

	// Register iterates over all routes added to the RouteRegister
	// and add them to the `Router` pass as an parameter.
	Register(Router)
}

type RouteRegisterImpl struct {
	router Router
}

func (rr *RouteRegisterImpl) Register(router Router) {
	rr.router = router
}

func (rr *RouteRegisterImpl) Get(pattern string, handler web.HandlerFunc) {
	rr.router.Add(http.MethodGet, pattern, handler)
}

func (rr *RouteRegisterImpl) Post(pattern string, handler web.HandlerFunc) {
	rr.router.Add(http.MethodPost, pattern, handler)
}

func (rr *RouteRegisterImpl) Delete(pattern string, handler web.HandlerFunc) {
	rr.router.Add(http.MethodDelete, pattern, handler)
}

func (rr *RouteRegisterImpl) Put(pattern string, handler web.HandlerFunc) {
	rr.router.Add(http.MethodPut, pattern, handler)
}

func (rr *RouteRegisterImpl) Patch(pattern string, handler web.HandlerFunc) {
	rr.router.Add(http.MethodPatch, pattern, handler)
}
