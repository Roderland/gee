package gee_cache

import "sync"

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type CallGroup struct {
	mu sync.Mutex
	m  map[string]*call
}

func (g *CallGroup) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}

	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	delete(g.m, key)
	return c.val, c.err
}
