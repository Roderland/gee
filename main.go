package main

import (
	"fmt"
	"gee_web"
	"net/http"
)

func main() {
	engine := gee_web.New()

	engine.Use(gee_web.Logger())

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

	group_test(engine)

	v3 := engine.NewGroup("/v3")
	v3.Use(gee_web.Failed())
	{
		v3.GET("/hello/:name", func(c *gee_web.Context) {
			c.String(200, "hello")
		})
	}

	engine.Run(":8080")
}

func group_test(engine *gee_web.Engine) {
	engine.GET("/index", func(c *gee_web.Context) {
		c.HTML(200, "<h1>Index Page</h1>")
	})

	v1 := engine.NewGroup("/v1")
	{
		v1.GET("/hello/:name", func(c *gee_web.Context) {
			c.HTML(200, fmt.Sprintf("<h1>hello %s<h1>", c.Param("name")))
		})
	}

	v2 := engine.NewGroup("/v2")
	{
		v2.POST("/login", func(c *gee_web.Context) {
			c.JSON(200, map[string]interface{}{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
}
