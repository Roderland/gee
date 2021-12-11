package main

import (
	"flag"
	"fmt"
	"gee_cache"
	"gee_web"
	"html/template"
	"log"
	"net/http"
	"time"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func createGroup() *gee_cache.Group {
	return gee_cache.NewGroup("scores", 2<<10, gee_cache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

func startCacheServer(addr string, addrs []string, gee *gee_cache.Group) {
	peers := gee_cache.NewHTTPPool(addr)
	peers.Set(addrs...)
	gee.RegisterPeers(peers)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

func startAPIServer(apiAddr string, gee *gee_cache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := gee.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())

		}))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))

}

func main() {
	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "Geecache server port")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	gee := createGroup()
	if api {
		go startAPIServer(apiAddr, gee)
	}
	startCacheServer(addrMap[port], []string(addrs), gee)
}

func mainWeb() {
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
