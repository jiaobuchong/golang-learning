package main

import "fmt"
import "geometry/computer"

func main() {

	var spec computer.Spec

	// 访问导出字段
	spec.Maker = "apple"
	spec.Price = 50000

	fmt.Println(spec)

}
