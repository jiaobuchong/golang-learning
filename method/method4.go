package main

import "fmt"

type rectangle1 struct {
	length int
	width  int
}

// 指针参数 的函数
func perimeter(r *rectangle1) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

// 指针接收器的方法，相当于 java 类的方法
func (r *rectangle1) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

// 在方法中使用「指针接收器」与在函数中使用「指针参数」
func main() {
	// r 是一个具体的值类型
	r := rectangle1{
		length: 10,
		width:  5,
	}
	p := &r //pointer to r
	perimeter(p)
	p.perimeter()

	/*
	   cannot use r (type rectangle) as type *rectangle in argument to perimeter
	   perimeter(r) 中的参数类型是指针类型，r 是值类型，传进去是不匹配的
	*/
	//perimeter(r)

	// r 是一个对象引用，通过这个对象引用可以直接调用 perimeter() 方法，底层还是 &r
	// &r 表示 r 的地址，指向了 r 的具体内存值数据
	r.perimeter() //使用值来调用指针接收器，也就是 (&r).perimeter()

}
