package gee_web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HandlerFunc func(*Context)

type Context struct {
	w      http.ResponseWriter
	r      *http.Request
	path   string
	method string
	params map[string]string
	code   int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		w:      w,
		r:      r,
		path:   r.URL.Path,
		method: r.Method,
	}
}

func (c *Context) Param(key string) string {
	return c.params[key]
}

func (c *Context) PostForm(key string) string {
	return c.r.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.r.URL.Query().Get(key)
}

func (c *Context) SetStatus(code int) {
	c.code = code
	c.w.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.w.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatus(code)
	c.w.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(code)
	encoder := json.NewEncoder(c.w)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.w, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.SetStatus(code)
	c.w.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatus(code)
	c.w.Write([]byte(html))
}
