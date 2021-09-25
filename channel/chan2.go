package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	// 写入信道
	done <- true
}

func main() {
	done := make(chan bool)
	go hello(done)
	// 从信道读取数据，这一行代码会发生阻塞，除非有协程向 done 写数据，否则程序不会跳到下一行代码
	// <-done 这行代码通过协程（译注：原文笔误，信道）done 接收数据，但并没有使用数据或者把数据存储到变量中。
	// 这完全是合法的。
	<-done
	fmt.Println("main function")
}
