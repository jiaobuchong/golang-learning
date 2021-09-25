package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

// 定义结构体类型

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

// 由于 Permanent 和 Contract 都声明了该方法，因此它们都实现了 SalaryCalculator 接口
//salary of permanent employee is sum of basic pay and pf
// 实现了接口 SalaryCalculator

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
// 实现了接口 SalaryCalculator

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

/*
total expense is calculated by iterating though the SalaryCalculator slice and summing
the salaries of the individual employees
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

}
