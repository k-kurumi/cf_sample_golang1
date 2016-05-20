package main

import (
	"github.com/kataras/iris"
	"os"
)

func main() {
	iris.Get("/hi", func(ctx *iris.Context) {
		page := map[string]interface{}{"Name": "Iris"}
		ctx.Render("hi.html", page)
	})
	iris.Listen(":" + os.Getenv("PORT"))
}
