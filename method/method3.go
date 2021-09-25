package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

// 接收一个值参数
func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

// 接收一个值接收器
func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

// 在方法中使用「值接收器」与在函数中使用「值参数」
func main() {

	// 当一个函数有一个值参数，它只能接受一个值参数。
	// 当一个方法有一个值接收器，它可以接受值接收器和指针接收器。
	r := rectangle{
		length: 10,
		width:  5,
	}

	// 通过值参数调用 area(r) 这个函数
	// 通过值值接收器来调用 r.area()
	area(r)
	r.area()

	// 指向 r 的指针 p
	p := &r
	/*
	   compilation error, cannot use p (type *rectangle) as type rectangle
	   in argument to area
	*/
	//area(p)

	p.area() // 通过指针调用值接收器

}
