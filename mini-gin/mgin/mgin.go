package mgin

import (
	"html/template"
	"net/http"
	"path"
	"strings"
)

type RouterGroup struct {
	prefix      string
	middlewares []HandleFunc // support middleware
	parent      *RouterGroup // support nesting
	engine      *Engine      // all groups share a Engine instance
}

// HandleFunc HandlerFunc defines the request handler
type HandleFunc func(c *Context)

// Engine implement the interface of ServerHttp
type Engine struct {
	*RouterGroup
	router        *router
	groups        []*RouterGroup     // store all groups
	htmlTemplates *template.Template // for html template
	funcMap       template.FuncMap   // for html render
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHtmlGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}

// New init mgin.Engin
func New() *Engine {
	currentEngine := &Engine{router: newRouter()}
	currentEngine.RouterGroup = &RouterGroup{engine: currentEngine}
	currentEngine.groups = []*RouterGroup{currentEngine.RouterGroup}
	return currentEngine
}

// gee.go
// Default use Logger() & Recovery middlewares
func Default() *Engine {
	engine := New()
	engine.Use(Recovery())
	return engine
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
	var middlewares []HandleFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(request.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(response, request)
	c.handlers = middlewares
	c.engine = engine
	engine.router.handle(c)
}

func (group *RouterGroup) Use(middlewares ...HandleFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// static handler
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandleFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	var handleFunc = func(c *Context) {
		file := c.Param("filepath")
		// check if file exist or we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
	return handleFunc
}

// save static file
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register get handler
	group.GET(urlPattern, handler)
}
