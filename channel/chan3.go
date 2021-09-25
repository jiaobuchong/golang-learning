package main

import (
	"fmt"
	"time"
)

func hello1(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	// 开启 hello1 协程
	go hello1(done)
	<-done
	fmt.Println("Main received data")
}
