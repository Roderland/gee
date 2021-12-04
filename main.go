package main

import (
	"fmt"
	"gee_web"
	"net/http"
)

func main() {
	engine := gee_web.New()
	engine.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})
	engine.Run(":8080")
}