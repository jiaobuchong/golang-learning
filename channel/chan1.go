package main

import "fmt"

// 信道可以想像成 Go 协程之间通信的管道。如同管道中的水会从一端流到另一端，
// 通过使用信道，数据也可以从一端发送，在另一端接收。
// 所有信道都关联了一个类型。信道只能运输这种类型的数据，而运输其他类型的数据都是非法的。
func main() {
	// 程序中 a 是一个 int 类型的信道
	// 信道好比一个阻塞队列，类似 SynchroQueue
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		// 用 make 来定义信道
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}

}
