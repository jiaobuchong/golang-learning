package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Int 和 String 可以帮助我们分别取出 reflect.Value 作为 int64 和 string
	//取出 reflect.Value，并转换为 int64
	//取出 reflect.Value 并将其转换为 string
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)

}
