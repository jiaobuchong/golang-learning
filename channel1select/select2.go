package main

func main() {
	ch := make(chan string)
	select {
	// 没有 go 协程向该信道写入数据，select 语句会一直阻塞，导致死锁
	case <-ch:
	}
}
