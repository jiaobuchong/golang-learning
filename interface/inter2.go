package main

import (
	"fmt"
)

type Test interface {
	Tester()
}

type MyFloat float64

func (m MyFloat) Tester() {
	fmt.Println(m)
}

func describe(t Test) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func main() {
	var t Test
	f := MyFloat(89.7)

	// 可以把接口看作内部的一个元组 (type, value)。 type 是接口底层的具体类型（Concrete Type），而 value 是具体类型的值。
	// t 的具体类型是 MyFloat，t 的值是 89.7
	t = f
	describe(t)
	t.Tester()
}
