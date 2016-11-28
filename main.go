package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

// golangのライブラリはどうしてすぐにリリースを削除するのだろう

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, gooolang!!")
	})

	// 遅延して応答する
	e.GET("/delay/:second", getDelay)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

// e.GET("/delay/:second", getDelay)
func getDelay(c echo.Context) error {
	second, _ := strconv.Atoi(c.Param("second"))
	time.Sleep(time.Duration(second) * time.Second)
	return c.String(http.StatusOK, fmt.Sprintf("%d delay", second))
}
