package main

// select 语句用于在多个发送/接收信道操作中进行选择。
// select 语句会一直阻塞，直到发送/接收操作准备就绪。
// 如果有多个信道操作准备完毕，select 会随机地选取其中之一执行。该语法与 switch 类似，
// 所不同的是，这里的每个 case 语句都是信道操作。我们好好看一些代码来加深理解吧。

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	// select 会阻塞，除非其中有 case 准备就绪
	select {
		case s1 := <-output1:
			fmt.Println(s1)
		case s2 := <-output2:
			// 3 秒之后信道 s2 有数据然后终止执行
			fmt.Println(s2)
	}
}
