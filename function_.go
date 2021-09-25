package main

// 所有 go 程序都必须包含一个 main 函数，这个函数是程序的入口，main 函数应该放置于 main 包中
import "fmt"

// 参数名 类型
func calBill(price int, no int) int {
	var totol = price * no

	return totol

}

// 多返回值，牛逼啊
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// 多返回值
func rectProps1(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}

// 可变参数函数，可变参数函数的工作原理是把可变参数转换为一个新的切片
func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func change(s ...string) {
	s[0] = "Go"
	// 当新的元素被添加到切片时，会创建一个新的数组。现有数组的元素被复制到这个新数组中，
	// 并返回这个新数组的新切片引用。现在新切片的容量是旧切片的两倍。
	s = append(s, "playground")
	fmt.Println("ddd: ", s)
}

func main() {
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	// _ 空白符，表示任何类型的任何值
	// 计算面积，不关心周长的结果
	area1, _ := rectProps(10.8, 5.6) // 返回值周长被丢弃
	fmt.Printf("Area %f ", area1)

	welcome := []string{"hello", "world"}

	// 如果使用了 ... ，welcome 切片本身会作为参数直接传入，不需要再创建一个新的切片
	change(welcome...)
	fmt.Println(welcome)
}
