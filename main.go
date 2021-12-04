package main

import (
	"gee_web"
)

func main() {
	engine := gee_web.New()
	engine.GET("/", func(c *gee_web.Context) {
		c.String(200, "hello %s", c.Query("name"))
	})
	engine.POST("/login", func (c *gee_web.Context) {
		c.JSON(200, map[string]interface{}{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	engine.Run(":8080")
}