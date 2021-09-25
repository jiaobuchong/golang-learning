package main

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	sendch := make(chan<- int)
	go sendData(sendch)
	// 这个信道是 send only，唯送信道，
	// 试图通过唯送信道接收数据，编译器报错
	// fmt.Println(<-sendch)

	// 只不过一个不能读取数据的唯送信道究竟有什么意义呢？
	// 这就需要用到信道转换（Channel Conversion）了。把一个双向信道转换成唯送信道（Send Only）
	// 或者唯收（Receive Only）
	// 信道都是行得通的，但是反过来就不行。
}
