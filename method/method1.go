package main

import (
	"fmt"
)

type Employee struct {
	name string
	age  int
}

/*
使用值接收器的方法。
*/
func (e Employee) changeName(newName string) {
	e.name = newName
}

/*
使用指针接收器的方法。
指针接收器的方法内部的改变对于调用者是可见的，
然而值接收器的情况不是这样的
*/
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	// 值接收器对调用者不可见，所以这里没有改变成功
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	// 指针接收器的方式可以改变值
	e.changeAge(43) // 这个可以自动解释成 &e，如下：
	(&e).changeAge(51)

	fmt.Printf("\nEmployee age after change: %d", e.age)

	/*
		那么什么时候使用指针接收器，什么时候使用值接收器？
		一般来说，指针接收器可以使用在：对方法内部的接收器所做的改变应该对调用者可见时。

		指针接收器也可以被使用在如下场景：当拷贝一个结构体的代价过于昂贵时。考虑下一个结构体有很多的字段。
		在方法内使用这个结构体做为值接收器需要拷贝整个结构体，
		这是很昂贵的。在这种情况下使用指针接收器，结构体不会被拷贝，只会传递一个指针到方法内部使用。

		在其他的所有情况，值接收器都可以被使用。
	*/
}
