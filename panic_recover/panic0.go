package main

import "fmt"

// 需要注意的是，你应该尽可能地使用错误，而不是使用 panic 和 recover。
// 只有当程序不能继续运行的时候，才应该使用 panic 和 recover 机制。
// 调用了 panic 之后，程序会终止，类似 Java System.exit()

/*
panic 的两个合理的调用：
1. 发生了一个不能恢复的错误，此时程序不能继续运行
2. 发生了一个编程上的错误
 */

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		// 发送 panic 程序就会终止
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}