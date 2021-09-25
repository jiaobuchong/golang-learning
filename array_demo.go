package main

import (
	"fmt"
)

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	a := [6]int{12, 78, 50, 3532, 464, 565} // short hand declaration to create array
	b := a

	// 数组的引用是值引用，改变 copy 中的值，不会改变原数组中的值
	b[0] = 123

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(len(a))

	sum := 0
	// 数组遍历，i 是索引
	for i, v := range a { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println(sum)

	// 切片，对切片所做的修改都会反映在底层数组中
	dslice := a[2:4]
	fmt.Println("array before ", a)

	for j := range dslice {
		dslice[j]++
	}

	fmt.Println("array after ", a)
}
