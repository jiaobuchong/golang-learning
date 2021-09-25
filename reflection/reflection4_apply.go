package main

import (
	"fmt"
	"reflect"
)

type order4 struct {
	ordId      int
	customerId int
}

type employee4 struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery4(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		// 获取结构体名称
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		// 获取结构体 q 的值
		v := reflect.ValueOf(q)
		// 获取字段个数
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}

func main() {
	o := order4{
		ordId:      456,
		customerId: 56,
	}
	createQuery4(o)

	e := employee4{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery4(e)
	i := 90
	createQuery4(i)

}