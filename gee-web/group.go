package gee_web

import (
	"net/http"
	"path"
)

type Group struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *Group
	engine      *Engine
}

func (g *Group) NewGroup(prefix string) *Group {
	engine := g.engine
	newGroup := &Group{
		prefix: prefix,
		parent: g,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (g *Group) Use(middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *Group) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *Group) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

func (g *Group) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}

// create static handler
func (g *Group) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(g.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.SetStatus(404)
			return
		}
		fileServer.ServeHTTP(c.w, c.r)
	}
}

// serve static files
func (g *Group) Static(relativePath string, root string) {
	handler := g.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	g.GET(urlPattern, handler)
}
