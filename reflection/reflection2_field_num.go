package main

import (
	"fmt"
	"reflect"
)

type order2 struct {
	ordId      int
	customerId int
}

func createQuery2(q interface{}) {
	//因为 NumField 方法只能在结构体上使用，这里需要判断 q 的类别是 struct
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		// 获取结构体实际类型
		t := reflect.TypeOf(q)
		// 反射获取值
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		//NumField() 方法返回结构体中字段的数量
		//Field(i int) 方法返回字段 i 的 reflect.Value
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T Field name: %s value:%v\n", i, v.Field(i),
				t.Field(i).Name,
				v.Field(i))
		}
	}

}
func main() {
	o := order2{
		ordId:      456,
		customerId: 56,
	}
	createQuery2(o)
}