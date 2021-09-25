package main

import "hello/oop/employee"

func main() {
	// 零值变量，不会导致空指针异常
	// 通过 oop2 中的方式来避免创建不可用的 employee 结构体变量
	var e employee.Employee
	e.LeavesRemaining()
}