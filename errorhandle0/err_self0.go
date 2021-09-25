package main

import (
	"errors"
	"fmt"
	"math"
)

// 通过 New 来自定义错误
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		// 打印的是 error 接口的 Error 方法返回的内容
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}
