package main

import (
	"fmt"
)

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string

	// 匿名字段
	address
}

func main() {
	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "California",
		},
	}

	//属于结构体的匿名字段的方法可以被直接调用，就好像这些方法是属于定义了匿名字段的结构体一样。
	p.fullAddress() //访问 address 结构体的 fullAddress 方法
}
