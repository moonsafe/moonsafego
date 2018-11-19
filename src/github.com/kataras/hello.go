package main

import (
	"github.com/kataras/iris"
)

type moon struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Poster string `json:"poster"`
}

func main() {
	m := new(moon)
	m.ID = 1
	m.Name = "jack"
	m.Year = 1992
	m.Genre = "man"
	m.Poster = "hello world!"

	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", m)
		ctx.View("hello.html")
	})
	app.Run(iris.Addr(":8888"))

}
