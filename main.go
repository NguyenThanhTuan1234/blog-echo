package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.File("/index.html", "public/index.html")
	e.Static("/*", "public")
	e.Static("/css/*", "public/css")
	e.Static("/js/*", "public/js")
	e.Static("/img/*", "public/img")
	e.Logger.Fatal(e.Start(":1323"))
}
