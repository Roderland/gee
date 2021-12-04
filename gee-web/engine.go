package gee_web

import (
	"net/http"
)

type Engine struct {
	*Group
	router *router
	groups []*Group
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.Group = &Group{engine: engine}
	engine.groups = []*Group{engine.Group}
	return engine
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := newContext(w, r)
	e.router.handle(context)
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	e.addRoute("GET", pattern, handlerFunc)
}

func (e *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	e.addRoute("POST", pattern, handlerFunc)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
