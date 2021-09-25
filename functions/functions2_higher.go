package main

import (
	"fmt"
)

// 满足下列条件之一函数都叫高阶函数（Hiher-order Function）：
// 1. 接收一个或多个函数作为参数
// 2. 返回值是一个函数

// 函数作为参数
func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

// 在其他函数中返回函数，返回类型 func(a, b int) int
func simple1() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func main() {
	f := func(a, b int) int {
		return a + b
	}
	simple(f)

	s := simple1()
	fmt.Println(s(45, 45))
}
