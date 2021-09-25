package main

import "fmt"

func sendData1(sendch chan<- int) {
	// channel conversion
	// 参数 sendch chan<- int 把 cha1 转换为一个唯送信道
	//于是该信道在 sendData 协程里是一个唯送信道，而在 Go 主协程里是一个双向信道。
	//该程序最终打印输出 10。
	sendch <- 10
}

func main() {
	cha1 := make(chan int)
	go sendData1(cha1)
	fmt.Println(<-cha1)
}
