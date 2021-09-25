package main

// 工作池就是一组等待任务分配的线程。一旦完成了所分配的任务，这些线程可继续等待任务的分配。
// 使用缓冲信道来实现工作池

/*
我们工作池的核心功能如下：

1. 创建一个 Go 协程池，监听一个等待作业分配的输入型缓冲信道。
2. 将作业添加到该输入型缓冲信道中。
3. 作业完成后，再将结果写入一个输出型缓冲信道。
4. 从输出型缓冲信道读取并打印结果。
*/

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job Job
	// 计算的结果
	sumofdigits int
}

// 接收作业缓冲信道
var jobs = make(chan Job, 10)

// 写入结果缓冲信道
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

// 创建工作协程
func worker(wg *sync.WaitGroup) {
	// 读取任务，执行任务
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	// 完成 job 之后
	wg.Done()
}

//创建了一个 Go 协程的工作池
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		// 创建协程
		go worker(&wg)
	}
	// 创建工作协程以后，Wait() 等待所有的 Go 协程执行完毕
	// 这里阻塞住了主线程
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	// 将任务放入队列
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	// 关闭信道，表示不再往这个信道写入数据了
	// golang 里的 close 只是用于通知信道的接收方，所有数据都已经发送完毕，信道没有真正关闭。
	// 若用 for range 接收数据时，对于关闭了的信道，会接收完剩下的有效数据，并退出循环。
	// 如果没有 close 提示数据发送完毕的话，
	// for range 会接收完剩下所有有效数据后发生阻塞。
	// 所以接收方 worker 是可以把 jobs 剩下的数据取走的。后面垃圾收集器会自动回收掉该信道的内存。
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	noOfJobs := 100
	// 将任务放入信道 jobs
	go allocate(noOfJobs)
	done := make(chan bool)
	// 从结果信道 results 获取结果
	go result(done)    // 这行如果放在 createWorkerPool 之后会死锁，因为 createWorkerPool 往 results 信道放任务，
	                   // 没协程消费阻塞导致死锁
	noOfWorkers := 10

	// 创建 10 个协程处理任务
	// createWorkerPool 里面的等待只是用来等待一些列子协程完成各自任务，完成任务并投放到results通道
	createWorkerPool(noOfWorkers)
	fmt.Println("create goroutine ready")

	// 这个是用来等待消费者，也就是等待 go result 协程去完成任务以后才可以退出程序，否则可能出现提前退出的情况，
	// 也就是还没有消费完，go 主进程已经结束（守护进程）
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
