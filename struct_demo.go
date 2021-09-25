package main

import "fmt"

//命名结构体
type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

// 匿名结构体
var employee struct {
	firstName, lastName string
	age                 int
}

func main() {

	//creating structure using field names
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}
	emp12 := Employee{}
	fmt.Println("zero value: ", emp12.lastName == "")

	//creating structure without using field names
	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	// 匿名结构体
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)

	// 结构体默认值，string 类型是 ""，int 类型是 0 值
	var emp6 Employee
	fmt.Println("employee 4", emp6)
}
