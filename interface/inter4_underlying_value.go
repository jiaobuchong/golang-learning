package main

import (
	"fmt"
)

// 类型断言用于提取接口的底层值
func assert(i interface{}) {
	// i.(int) 来提取 i 的底层 int 值
	s := i.(int) //get the underlying int value from i
	fmt.Println(s)
}

func assert1(i interface{}) {
	// 如果 i 不是 int 类型，ok 为 false
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func main() {
	//
	var s interface{} = 56
	assert(s)

	var s1 interface{} = "jack"
	assert1(s1)
}