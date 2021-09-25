package main

import "fmt"

type myInt int

// 方法接收器类型的定义和方法的定义应该在同一个包中
func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}
