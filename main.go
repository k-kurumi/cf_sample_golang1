package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
)

// 並列数の制限
// http://qiita.com/ruiu/items/d4ef16981b870a1b85af

// このWaitGroup方式ではメモリが溢れた
// http://deeeet.com/writing/2014/07/30/golang-parallel-by-cpu/

var cpus int

func main() {
	cpus := runtime.NumCPU()
	e := echo.New()
	e.Get("/", func(c echo.Context) error {

		ch := make(chan bool, cpus)
		for i := 1; i <= 500; i++ {
			ch <- true
			go func(i int) {
				defer func() { <-ch }()
				fmt.Println("mylog", i)
			}(i)
		}

		return c.String(http.StatusOK, "Hellloooo, world")
	})

	e.Run(standard.New(":" + os.Getenv("PORT")))
}
