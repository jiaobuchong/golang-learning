package main

import (
	"fmt"
	"reflect"
)

type order1 struct {
	ordId      int
	customerId int
}

func createQuery1(q interface{}) {
	//reflect.Type 表示 interface{} 的具体类型，而 reflect.Value 表示它的具体值
	//reflect.TypeOf() 和 reflect.ValueOf() 两个函数可以分别返回 reflect.Type 和 reflect.Value
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)

	//Kind 和 Type 的类型可能看起来很相似，但在下面程序中，可以很清楚地看出它们的不同之处
	k := t.Kind()
	fmt.Println("Kind ", k)
	//Type 表示 interface{} 的实际类型（在这里是 main.Order)，而 Kind 表示该类型的特定类别（在这里是 struct）

	fmt.Println("Type ", t)
	fmt.Println("Value ", v)

}

func main() {
	o := order1{
		ordId:      456,
		customerId: 56,
	}
	createQuery1(o)

}