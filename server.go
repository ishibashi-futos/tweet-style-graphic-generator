package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Twitter風HTML参考
// https://3owebcreate.com/web/coding/chat_twitter_css

func serve() {
	t := &Template{
		templates: template.Must(template.ParseGlob("static/*.html")),
	}
	e := echo.New()
	e.Static("/", "./static/")
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		var data interface{}
		return c.Render(http.StatusOK, "index", data)
	})

	e.Start(":3000")
}
