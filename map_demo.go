package main

import "fmt"

func main()  {

	// 通过向 make 函数传入键和值的类型，
	// 可以创建 map。make(map[type of key]type of value) 是创建 map 的语法。

	var personSalary1 map[string]int
	if personSalary1 == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary1 = make(map[string]int)
	}

	// 健是 string 类型，值是 int 类型
	personSalary := make(map[string]int)
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000

	fmt.Println(personSalary)

	// 通过 ok 判断这个元素是否存在，true 表示存在
	value, ok := personSalary["mike"]
	if ok {
		fmt.Println("value: ", value, "ok", ok)
	}

	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}





}
