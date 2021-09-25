package main

import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	ch := make(chan int)
	go producer(ch)

	// 死循环
	for {
		v, ok := <-ch
		// 使用 ok 检查信道是否关闭
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}
