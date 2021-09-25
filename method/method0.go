package main

import "fmt"

type Employee1 struct {
	name     string
	salary   int
	currency string
}

/*
  方法
  displaySalary() 方法将 Employee 做为接收器类型
*/
func (e Employee1) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

/*
displaySalary()方法被转化为一个函数，把 Employee 当做参数传入。
*/
func displaySalary(e Employee1) {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

// 方法的使用
/**
go 不支持类，基于类型的方法是一种实现和类相似行为的途径
 */
func main() {

	emp1 := Employee1 {
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}

	// 这个看上去类似 java 的类方法
	emp1.displaySalary() // 调用 Employee 类型的 displaySalary() 方法

}
