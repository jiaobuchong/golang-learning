package main

import "fmt"

type areaError struct {
	err    string  //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

// 实现 error 接口，并给错误类型添加了两个方法，提供了更多的错误信息
func (e *areaError) Error() string {
	return e.err
}

// 两个方法都提供了关于错误的更多信息，在这里，它提示我们计算面积失败的原因（长度为负数或者宽度为负数）。
// 于是我们就有了两个错误类型结构体的方法，来提供更多的错误信息。
func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError{err, length, width}
	}
	return length * width, nil
}

func main() {
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		// 断言 *areaError 类型
		// 类似 java 的强制类型转换
		if err, ok := err.(*areaError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)

			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)

			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)
}
