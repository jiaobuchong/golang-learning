package main

import (
	"fmt"
)

// 没有包含方法的接口是空接口
// 这个函数可以接收任何类型
func describe1(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Hello World"

	// 所有类型都实现了空接口
	describe1(s)
	i := 55
	describe1(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe1(strt)
}
