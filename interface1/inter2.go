package main

import (
	"fmt"
)

type SalaryCalculator1 interface {
	DisplaySalary()
}

type LeaveCalculator1 interface {
	CalculateLeavesLeft() int
}

//创建了一个新的接口 EmployeeOperations，它嵌套了两个接口：SalaryCalculator 和 LeaveCalculator。
//如果一个类型定义了 SalaryCalculator 和 LeaveCalculator 接口里包含的方法，
//我们就称该类型实现了 EmployeeOperations 接口。


type EmployeeOperations interface {
	// 嵌套接口
	SalaryCalculator1
	LeaveCalculator1
}

type Employee1 struct {
	firstName string
	lastName string
	basicPay int
	pf int
	totalLeaves int
	leavesTaken int
}

//由于 Employee 结构体定义了 DisplaySalary 和 CalculateLeavesLeft 方法，
//因此它实现了接口 EmployeeOperations。

func (e Employee1) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee1) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

// 尽管 Go 语言没有提供继承机制，但可以通过嵌套其他的接口，创建一个新接口。
func main() {
	e := Employee1 {
		firstName: "Naveen",
		lastName: "Ramanathan",
		basicPay: 5000,
		pf: 200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
}