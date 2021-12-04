package main

import (
	"gee_web"
	"net/http"
)

func main() {
	engine := gee_web.New()
	engine.GET("/", func(c *gee_web.Context) {
		c.String(200, "hello %s", c.Query("name"))
	})
	engine.POST("/login", func(c *gee_web.Context) {
		c.JSON(200, map[string]interface{}{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	engine.GET("/hello/:name", func(c *gee_web.Context) {
		c.String(http.StatusOK, "hello %s\n", c.Param("name"))
	})
	engine.GET("/assets/*filepath", func(c *gee_web.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{"filepath": c.Param("filepath")})
	})
	engine.Run(":8080")
}
