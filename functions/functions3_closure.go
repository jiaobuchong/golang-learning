package main

import "fmt"

// 闭包
//闭包（Closure）是匿名函数的一个特例。
//当一个匿名函数所访问的变量定义在函数体的外部时，就称这样的匿名函数为闭包。


//每一个闭包都会绑定一个它自己的外围变量（Surrounding Variable）。
func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	a := 5
	func() {
		fmt.Println("a =", a)
	}()

	// 函数 appendStr 返回了一个闭包。这个闭包绑定了变量 t。
	//每一个闭包都会绑定一个它自己的外围变量（Surrounding Variable）。
	// c 和 b 都是闭包，这个闭包绑定了变量 t，是一个独立不会改变的字段
	c := appendStr()
	b := appendStr()

	// 用参数 World 调用了 c。现在 c 中 t 值变为了 Hello World。
	fmt.Println(c("World"))
	//用参数 Everyone 调用了 b。由于 b 绑定了自己的变量 t，
	//因此 b 中的 t 还是等于初始值 Hello。于是该函数调用之后，b 中的 t 变为了 Hello Everyone。
	fmt.Println(b("Everyone"))

	fmt.Println(c("Gopher"))
	fmt.Println(b("!"))
}
