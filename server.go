package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// Twitter風HTML参考
// https://3owebcreate.com/web/coding/chat_twitter_css

func serve() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})

	e.Start(":3000")
}
