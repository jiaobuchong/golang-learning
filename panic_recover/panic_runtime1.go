package main

import (
	"fmt"
	"runtime/debug"
)

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		// 打印出堆栈跟踪
		debug.PrintStack()
	}
}

func a2() {
	// 从 panic 中恢复
	defer r()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

func main() {
	a2()
	fmt.Println("normally returned from main")
}
