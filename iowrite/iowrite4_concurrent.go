package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

// 并发写入文件
func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	f, err := os.Create("iowrite/concurrent.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 然后从 channel 中读取随机数并且写到文件中
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	//通过往 done 这个 channel 中写入 true 来通知任务已完成
	done <- true
}

func main() {
	// 创建写入和读取数据的 channel
	data := make(chan int)
	//创建 done 这个 channel
	//此 channel 用于消费者 goroutinue 完成任务之后通知 main 函数
	done := make(chan bool)

	//创建 Waitgroup 的实例 wg，用于等待所有生产随机数的 goroutine 完成任务
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func() {
		//waitgroup 的 wait() 方法等待所有的 goroutines 完成随机数的生成。
		//然后关闭 channel。
		wg.Wait()
		close(data)
	}()

	//将 true 写入 done 这个 channel 中，这个时候 main 函数解除阻塞
	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}