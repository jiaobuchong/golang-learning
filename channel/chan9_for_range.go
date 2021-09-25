package main

import (
	"fmt"
)

func producer1(chnl chan int) {
	for i := 0; i < 10; i++ {
		// 往信道写入数据
		chnl <- i
	}
	close(chnl)
}

// 关于信道还有一些其他的概念，比如缓冲信道（Buffered Channel）、工作池（Worker Pool）和 select。
func main() {
	ch := make(chan int)
	go producer1(ch)
	// for range 循环从信道 ch 接收数据，直到该信道关闭。一旦关闭了 ch，循环会自动结束。
	for v := range ch {
		fmt.Println("Received ", v)
	}
}
