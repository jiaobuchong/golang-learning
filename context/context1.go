package main

import (
	"context"
	"fmt"
	"time"
)

// 更复杂的场景如何做并发控制呢？比如子协程中开启了新的子协程，或者需要同时控制多个子协程。
// 这种场景下，select+chan 的方式就显得力不从心了。

// Go 语言提供了 Context 标准库可以解决这类场景的问题，Context 的作用和它的名字很像，
// 上下文，即子协程的下上文。Context 有两个主要的功能：
// 	1. 通知子协程退出（正常退出，超时退出等）；
//  2. 传递必要的参数。

//https://geektutu.com/post/quick-go-context.html

func reqTask1(ctx context.Context, name string) {
	for {
		select {
		// 在子协程中，使用 select 调用 <-ctx.Done() 判断是否需要退出
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// 创建可取消的 context 对象
	// context.Backgroud() 创建根 Context，通常在 main 函数、初始化和测试代码中创建，作为顶层 Context
	//context.WithCancel(parent) 创建可取消的子 Context，同时返回函数 cancel
	ctx, cancel := context.WithCancel(context.Background())

	// 控制多个协程
	go reqTask1(ctx, "worker1")
	go reqTask1(ctx, "worker2")

	time.Sleep(3 * time.Second)

	// 主协程中，调用 cancel() 函数通知子协程退出
	// 调用 cancel() 函数后该 Context 控制的所有子协程都会退出
	cancel()
	time.Sleep(3 * time.Second)
}