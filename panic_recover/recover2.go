package main

import (
	"fmt"
	"time"
)

func main() {

	// 这了就没法恢复了
	/*defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			//log.Error(nil, "panic", zap.Reflect("err", e))
		}
	}()
*/
	// 调用协程
	go func() {
		// 出现 panic 之后还可以恢复执行 到 	fmt.Println("recover from bala")
		defer func() {
			if e := recover(); e != nil {
				fmt.Println(e)
				//log.Error(nil, "panic", zap.Reflect("err", e))
			}
		}()
		for {
			time.Sleep(3 * time.Second)
			panic("oh! B panicked")
			//arr := []int{1, 3, 4}
			//fmt.Println(arr[3])
		}
	}()

	time.Sleep(10 * time.Second)

	fmt.Println("recover from bala")

}
