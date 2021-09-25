package main

// 无缓冲信道的发送和接收过程是阻塞的
// 本节主要是 buffered channel

// 可以创建一个有缓冲（Buffer）的信道。只在缓冲已满的情况，才会阻塞向缓冲信道（Buffered Channel）发送数据。
// 同样，只有在缓冲为空的时候，才会阻塞从缓冲信道接收数据。

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	// 不像无缓冲的 channel，这里不会阻塞
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
