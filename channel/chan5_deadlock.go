package main

func main() {
	ch := make(chan int)
	// 往信道写入数据，但没有协程消费，死锁
	ch <- 5
}
