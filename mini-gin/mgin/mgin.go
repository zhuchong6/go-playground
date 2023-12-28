package mgin

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler
type HandleFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServerHttp
type Engine struct {
	router map[string]HandleFunc
}

// init mgin.Engin
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

// add handler
func (engine *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	m := engine.router
	m[key] = handler
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
	key := request.Method + "-" + request.URL.Path
	if hanler, ok := engine.router[key]; ok {
		hanler(response, request)
	} else {
		fmt.Fprintf(response, "404 NOT FOUND: %s\n", request.URL)
	}
}
