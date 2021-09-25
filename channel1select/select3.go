package main

import (
	"fmt"
	"time"
)

func server3(ch chan string) {
	ch <- "from server1"
}
func server4(ch chan string) {
	ch <- "from server2"

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server3(output1)
	go server4(output2)
	time.Sleep(1 * time.Second)
	// 随机的选择一个 case 执行
	select {
		case s1 := <-output1:
			fmt.Println(s1)
		case s2 := <-output2:
			fmt.Println(s2)
	}
}
