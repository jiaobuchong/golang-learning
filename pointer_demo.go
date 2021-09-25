package main

import (
	"fmt"
)

func modify(arr *[3]int) {
	(*arr)[0] = 90
}

func modify1(sls []int) {
	sls[0] = 90
}

func main() {
	b := 255
	// a 指向了 b，a 存储的是 b 的地址
	var a *int = &b
	c := &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	// 获取指针指向的值
	fmt.Println("value of b is", *c)

	//
	*c++
	fmt.Println("value of b is", *c)

	// 向函数传递一个数组指针参数，并在函数内修改数组，尽管有效，但却不是 go 语言惯用的实现方式，
	// 最好使用切片来处理
	f := [3]int{89, 90, 91}
	// 切片会引用到原数组
	modify1(f[:])
	fmt.Println(f)

}
