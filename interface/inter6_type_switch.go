package main

import "fmt"

// 还可以将一个类型和接口相比较。如果一个类型实现了接口，那么该类型与其实现的接口就可以互相比较。

type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

//结构体 Person 实现了 Describer 接口

func (p Person) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType1(i interface{}) {
	// i.(type) 可以获取到具体的类型
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func main() {
	findType1("Naveen")
	p := Person{
		name: "Naveen R",
		age:  25,
	}
	findType1(p)
}
