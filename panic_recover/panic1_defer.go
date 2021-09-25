package main

import (
	"fmt"
)

func fullName1(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

//当函数发生 panic 时，它会终止运行，在执行完所有的延迟函数后，程序控制返回到该函数的调用方。
//这样的过程会一直持续下去，直到当前协程的所有函数都返回退出，然后程序会打印出 panic 信息，
//接着打印出堆栈跟踪，最后程序终止。

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName1(&firstName, nil)
	fmt.Println("returned normally from main")
}