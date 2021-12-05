package gee_web

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
