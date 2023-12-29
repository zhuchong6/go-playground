package mgin

import (
	"net/http"
)

// HandlerFunc defines the request handler
type HandleFunc func(c *Context)

// Engine implement the interface of ServerHttp
type Engine struct {
	router *router
}

// init mgin.Engin
func New() *Engine {
	return &Engine{router: newRouter()}
}

// add handler
func (engine *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// define GET
func (engine *Engine) GET(pattern string, handler HandleFunc) {
	engine.addRoute("GET", pattern, handler)
}

// define POST
func (engine *Engine) POST(pattern string, handler HandleFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// engine must implement ServeHTTP method, then you can http.ListenAndServe(addr, engine)
func (engine *Engine) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	c := newContext(response, request)
	engine.router.handle(c)
}
