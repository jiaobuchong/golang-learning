package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {

	// go hello() 启动了一个新的 Go 协程。现在 hello() 函数与 main()
	// 函数会并发地执行。主函数会运行在一个特有的 Go 协程上，它称为 Go 主协程（Main Goroutine）。

	// hello 协程
	go hello()

	//信道可用于在其他协程结束执行之前，阻塞 Go 主协程。

	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
