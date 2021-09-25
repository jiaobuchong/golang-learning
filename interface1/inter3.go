package main


import "fmt"

type Describer1 interface {
	Describe()
}

func main() {
	var d1 Describer1
	// 接口的零值是 nil。对于值为 nil 的接口，其底层值（Underlying Value）和具体类型（Concrete Type）都为 nil。
	if d1 == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
	}

	// d1 等于 nil，程序产生运行时错误 panic
	d1.Describe()
}