package main

import "hello/oop/employee"

func main() {
	//创建 employee 变量的唯一方法就是使用 New 函数。
	//虽然 Go 不支持类，但结构体能够很好地取代类，而以 New(parameters) 签名的方法可以替代构造器。
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
