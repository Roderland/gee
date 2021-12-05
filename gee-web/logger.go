package gee_web

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v\n", c.code, c.r.RequestURI, time.Since(t))
	}
}

func Failed() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v\n", c.code, c.r.RequestURI, time.Since(t))
	}
}
