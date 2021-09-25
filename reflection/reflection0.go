package main

import (
	"fmt"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name string
	id int
	address string
	salary int
	country string
}

func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}


// interface{} 可以接收任何结构体作为参数
func createQueryGeneric(q interface{}) {

	//reflect.Type 表示 interface{} 的具体类型，而 reflect.Value 表示它的具体值
	//reflect.TypeOf() 和 reflect.ValueOf() 两个函数可以分别返回 reflect.Type 和 reflect.Value


}



func main() {
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	fmt.Println(o)


	fmt.Println(createQuery(o))
}
