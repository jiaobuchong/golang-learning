package main

// 面向对象——继承
// Go 不支持继承，但它支持组合（Composition）。组合一般定义为“合并在一起”。
// 汽车就是一个关于组合的例子：一辆汽车由车轮、引擎和其他各种部件组合在一起。
// 通过组合来扩展功能

import (
	"fmt"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

// author 作为接受者类型
func (a author) fullName() string {
	// 类型 java String.format()
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	//嵌套的匿名字段 author
	//一旦结构体内嵌套了一个结构体字段，Go 可以使我们访问其嵌套的字段，好像这些字段属于外部结构体一样。
	// 组合结构体 author
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

type website struct {
	posts []post
}

func (w website) contents() {
	fmt.Println("Contents of Website\n")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}

	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{
		posts: []post{post1, post2, post3},
	}
	w.contents()
}
