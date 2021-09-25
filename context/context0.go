package main

import (
	"fmt"
	"time"
)

var stop chan bool

func reqTask(name string) {
	for {
		select {
		case <-stop:
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

// 更复杂的场景如何做并发控制呢？比如子协程中开启了新的子协程，或者需要同时控制多个子协程。
// 这种场景下，select+chan 的方式就显得力不从心了。
func main() {
	stop = make(chan bool)
	go reqTask("worker1")
	time.Sleep(3 * time.Second)
	stop <- true
	time.Sleep(3 * time.Second)
}
