package mgin

import (
	"net/http"
)

type RouterGroup struct {
	prefix      string
	middlewares []HandleFunc // support middleware
	parent      *RouterGroup // support nesting
	engine      *Engine      // all groups share a Engine instance
}

// HandlerFunc defines the request handler
type HandleFunc func(c *Context)

// Engine implement the interface of ServerHttp
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup // store all groups
}

// init mgin.Engin
func New() *Engine {
	currentEngine := &Engine{router: newRouter()}
	currentEngine.RouterGroup = &RouterGroup{engine: currentEngine}
	currentEngine.groups = []*RouterGroup{currentEngine.RouterGroup}
	return currentEngine
}

// group is defined to create new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// add handler
func (group *RouterGroup) addRoute(method string, pattern string, handler HandleFunc) {
	pt := group.prefix + pattern
	// log.Printf("Route %4s - %4s", method, pt)
	group.engine.router.addRoute(method, pt, handler)
}

// define GET
func (group *RouterGroup) GET(pattern string, handler HandleFunc) {
	group.addRoute("GET", pattern, handler)
}

// define POST
func (group *RouterGroup) POST(pattern string, handler HandleFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
// engine implement ServeHTTP method, engine can be seen as a Handler interface
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// engine must implement ServeHTTP method, then you can http.ListenAndServe(addr, engine)
func (engine *Engine) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	c := newContext(response, request)
	engine.router.handle(c)
}
