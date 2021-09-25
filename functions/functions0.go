package main

//支持头等函数（First Class Function）的编程语言，
//可以把函数赋值给变量，
//也可以把函数作为其它函数的参数或者返回值。
//Go 语言支持头等函数的机制。

import (
	"fmt"
)

func main() {

	//会发现赋值给 a 的函数没有名称。由于没有名称，这类函数称为匿名函数（Anonymous Function）
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T", a)
	fmt.Println()

	//调用一个匿名函数，可以不用赋值给变量。
	//使用 () 可以立即调用该匿名函数
	func() {
		fmt.Println("hello world second class function")
	}()

	// 向匿名函数传参
	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")
}