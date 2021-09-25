package main

//defer 语句的用途是：含有 defer 语句的函数，会在该函数将要返回之前，调用另一个函数。
//这个定义可能看起来很复杂，我们通过一个示例就很容易明白了。

import (
	"fmt"
)

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	// 类似切面，在函数执行 return 之前执行
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

func main() {
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)
}
