package main

import (
	"fmt"
)

func recoverName() {
	// 使用了 recover() 来停止 panic 续发事件
	// 在执行完 recover() 之后，panic 会停止
	if r := recover(); r!= nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName2(firstName *string, lastName *string) {
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName2(&firstName, nil)
	fmt.Println("returned normally from main")
}