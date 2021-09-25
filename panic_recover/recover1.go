package main

import (
	"fmt"
	"time"
)

//recover 不能恢复一个不同协程的 panic
func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	go b()
	//b() 这个可以恢复
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
}

func main() {
	a()
	fmt.Println("normally returned from main")
}
