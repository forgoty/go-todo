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

	// Group allows you to pass a function that can add multiple routes
	// with a shared prefix route.
	Group(string, func(RouteRegister))

	// Register iterates over all routes added to the RouteRegister
	// and add them to the `Router` pass as an parameter.
	Register(Router)
}

type RouteRegisterImpl struct {
	prefix string
	groups []*RouteRegisterImpl
	routes []route
}

type route struct {
	method  string
	pattern string
	handler web.HandlerFunc
}

// NewRouteRegister creates a new RouteRegister with all middlewares sent as params
func NewRouteRegister() *RouteRegisterImpl {
	return &RouteRegisterImpl{
		prefix: "",
		routes: []route{},
	}
}

func (rr *RouteRegisterImpl) Get(pattern string, handler web.HandlerFunc) {
	rr.route(http.MethodGet, pattern, handler)
}

func (rr *RouteRegisterImpl) Post(pattern string, handler web.HandlerFunc) {
	rr.route(http.MethodPost, pattern, handler)
}

func (rr *RouteRegisterImpl) Delete(pattern string, handler web.HandlerFunc) {
	rr.route(http.MethodDelete, pattern, handler)
}

func (rr *RouteRegisterImpl) Put(pattern string, handler web.HandlerFunc) {
	rr.route(http.MethodPut, pattern, handler)
}

func (rr *RouteRegisterImpl) Patch(pattern string, handler web.HandlerFunc) {
	rr.route(http.MethodPatch, pattern, handler)
}

func (rr *RouteRegisterImpl) Group(pattern string, fn func(rr RouteRegister)) {
	group := &RouteRegisterImpl{
		prefix: rr.prefix + pattern,
		routes: []route{},
	}

	fn(group)
	rr.groups = append(rr.groups, group)
}

func (rr *RouteRegisterImpl) Register(router Router) {
	for _, r := range rr.routes {
		router.Add(r.method, r.pattern, r.handler)
	}
	for _, g := range rr.groups {
		g.Register(router)
	}

}

func (rr *RouteRegisterImpl) route(method, pattern string, handler web.HandlerFunc) {
	fullPattern := rr.prefix + pattern
	for _, r := range rr.routes {
		if r.pattern == fullPattern && r.method == method {
			panic("cannot add duplicate route")
		}
	}

	rr.routes = append(rr.routes, route{
		method:  method,
		pattern: fullPattern,
		handler: handler,
	})
}
