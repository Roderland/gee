package main

import (
	"fmt"
	"gee_web"
	"html/template"
	"net/http"
	"time"
)

func main() {
	engine := gee_web.Default()

	engine.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	engine.LoadHtmlGlob("templates/*")
	engine.Static("/assets", "./static")
	stu1 := &student{"wg1", 18}
	stu2 := &student{"wg2", 20}

	engine.GET("/", func(c *gee_web.Context) {
		c.HTML(200, "css.tmpl", nil)
	})
	engine.GET("/students", func(c *gee_web.Context) {
		c.HTML(200, "arr.tmpl", gee_web.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	engine.GET("/date", func(c *gee_web.Context) {
		c.HTML(200, "custom_func.tmpl", gee_web.H{
			"title": "gee",
			"now":   time.Now(),
		})
	})

	engine.GET("/panic", func(c *gee_web.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	engine.Run(":8080")
}

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
