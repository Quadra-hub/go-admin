package main

import (
	"html/template"

	"github.com/Quadra-hub/go-admin/admin"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

type Photo struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

type Supply struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Photo `json:"photo"`
}

func main() {
	r := gin.Default()
	admin := admin.New()
	admin.Subscribe(Photo{})
	admin.Subscribe(Supply{})
	r.Static("/static", "./static")

	gv := goview.New(goview.Config{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/master",
		Funcs: template.FuncMap{
			"safeHTML": func(v string) template.HTML {
				return template.HTML(v)
			},
		},
		DisableCache: true,
	})

	goview.Use(gv)
	r.HTMLRender = ginview.Default()
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"tables": admin.GetTableNames(),
		})
	})

	r.Run()
}
