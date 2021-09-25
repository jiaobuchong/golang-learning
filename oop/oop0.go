package main

import "hello/oop/employee"

// 面向对象——封装
func main() {
	e := employee.Employee {
		FirstName: "Sam",
		LastName: "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
}