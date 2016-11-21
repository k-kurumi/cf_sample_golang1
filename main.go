package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// golangのライブラリはどうしてすぐにリリースを削除するのだろう

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, gooolang!!")
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
