package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"

	// 这里发生阻塞，deadlock
	ch <- "steve"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
