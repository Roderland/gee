package gee_web

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
 	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{make(map[string]HandlerFunc)}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handlerFunc, ok := e.router[key]; ok {
		handlerFunc(w, r)
	} else {
		fmt.Fprintf(w, "404 not found %s\n", r.URL)
	}
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
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