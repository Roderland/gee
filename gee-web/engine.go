package gee_web

import (
	"html/template"
	"net/http"
	"strings"
)

type Engine struct {
	*Group
	router        *router
	groups        []*Group
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.Group = &Group{engine: engine}
	engine.groups = []*Group{engine.Group}
	return engine
}

func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	context := newContext(w, r)
	context.handlers = middlewares
	context.engine = e
	e.router.handle(context)
}

func (e *Engine) SetFuncMap(funcMap template.FuncMap) {
	e.funcMap = funcMap
}

func (e *Engine) LoadHtmlGlob(pattern string) {
	e.htmlTemplates = template.Must(template.New("").Funcs(e.funcMap).ParseGlob(pattern))
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
