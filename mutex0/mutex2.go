package main

import (
	"fmt"
	"sync"
)

var x2 = 0

func increment2(wg *sync.WaitGroup, ch chan bool) {
	// 该缓冲信道用于保证只有一个协程访问增加 x2 的临界区
	//由于缓冲信道的容量为 1，所以任何其他协程试图写入该信道时，都会发生阻塞
	ch <- true
	x2 = x2 + 1
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	// 创建容量为1的缓冲信道
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment2(&w, ch)
	}
	w.Wait()

	fmt.Println("final value of x", x2)
}
