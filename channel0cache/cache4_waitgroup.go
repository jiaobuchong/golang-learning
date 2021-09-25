package main

//WaitGroup 用于等待一批 Go 协程执行结束。程序控制会一直阻塞，直到这些协程全部执行完毕。
//假设我们有 3 个并发执行的 Go 协程（由 Go 主协程生成）。Go 主协程需要等待这 3 个协程执行结束后，才会终止。
//这就可以用 WaitGroup 来实现。
import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	// 减少计数器
	wg.Done()
}

func main() {
	no := 3
	//
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		// 创建协程
		// 传递 wg 的地址是很重要的。如果没有传递 wg 的地址，那么每个 Go 协程将会得到一个 WaitGroup 值的拷贝，
		// 因而当它们执行结束时，main 函数并不会知道
		go process(i, &wg)
	}
	// 类似 java 的 CountDownLath
	// 计数器为 0 之后才会停止阻塞
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
