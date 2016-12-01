package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
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

	// pingする
	e.GET("/ping/:address", getPing)

	// 指定のHTTPステータスを返す
	e.GET("/status/:code", getStatus)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

// e.GET("/delay/:second", getDelay)
func getDelay(c echo.Context) error {
	second, _ := strconv.Atoi(c.Param("second"))
	time.Sleep(time.Duration(second) * time.Second)
	return c.String(http.StatusOK, fmt.Sprintf("delay: %d sec", second))
}

// 外部コマンド実行
func getPing(c echo.Context) error {
	address := c.Param("address")
	out, err := exec.Command("ping", "-c4", address).Output()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, string(out))
}

// 指定されたHTTPステータスを返す
func getStatus(c echo.Context) error {
	code, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(code, "")
}
