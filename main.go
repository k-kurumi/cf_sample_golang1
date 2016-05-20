package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/kataras/iris"
)

// 並列数の制限
// http://qiita.com/ruiu/items/d4ef16981b870a1b85af

// このWaitGroup方式ではメモリが溢れた
// http://deeeet.com/writing/2014/07/30/golang-parallel-by-cpu/

var cpus int

func main() {
	cpus := runtime.NumCPU()

	iris.Get("/", func(ctx *iris.Context) {

		c := make(chan bool, cpus)
		for i := 1; i <= 500; i++ {
			c <- true

			go func(i int) {
				defer func() { <-c }()
				fmt.Println("mylog", i)
			}(i)

		}

		ctx.Write("cpus:", cpus)
	})

	iris.Listen(":" + os.Getenv("PORT"))
}
