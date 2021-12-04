package gee_web

import "log"

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: map[string]HandlerFunc{}}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.method + "-" + c.path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(404, "404 not found %s\n", c.path)
	}
}
