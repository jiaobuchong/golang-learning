package main

import (
	"fmt"
)

func main() {
	// 3 是容量
	ch := make(chan string, 3)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println("capacity is", cap(ch))
	// len 是当前信道排队元素的个数
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch))
}
