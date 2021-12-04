package gee_web

import (
	"net/http"
)

type Engine struct {
 	router *router
}

func New() *Engine {
	return &Engine{newRouter()}
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