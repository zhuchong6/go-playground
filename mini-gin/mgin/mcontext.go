package mgin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// a util to process request info and response info
type Context struct {
	// net/http original objects
	Writer http.ResponseWriter
	Req    *http.Request

	// request info
	Path   string
	Method string
	Params map[string]string

	// response info
	StatusCode int
}

func (c *Context) Param(key string) string{
	value, _ := c.Params[key]
	return value
}

// init context
func newContext(response http.ResponseWriter, request *http.Request) *Context {
	context := &Context{
		Writer: response,
		Req:    request,
		Path:   request.URL.Path,
		Method: request.Method,
	}
	return context
}

// process postform param
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// process query param from url
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
