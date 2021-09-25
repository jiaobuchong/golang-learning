package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {

	array1 := [3]int{}

	array2 := [2]int{1, 2}
	array3 := [2]int{4, 5}

	// 数组属于基本类型，赋值属于值传递，
	// 数组在参数中传递，也是属于值传递
	array2 = array3

	array3[0] = 6

	fmt.Println(array1, array2, array3)
}

// 切片并不存储任何元素，只是对现有数组的引用（不是值拷贝，是指针）
func TestSlice(t *testing.T) {

	// 通过数组创建切片
	array1 := [3]int{1, 2, 3}
	//将数组下标从1处到下标2处的元素转换为一个切片(前闭后开)
	slice1 := array1[1:2]
	fmt.Println(slice1)
}

//https://www.cnblogs.com/sy270321/p/11388399.html
func TestSlice1(t *testing.T) {

	//下面代码直接出初始化一个切片 （这里大家有个疑问，我不管怎么看都觉得它是一个数组啊）
	//记住，再go语言中，区别一个变量是数组还是切片，就看有没有定义长度
	//有定义长度就是数组，如array1，没定义就是切片 如slice2
	//我们也通过fmt.Println(reflect.TypeOf(array1),reflect.TypeOf(slice2))
	//上面这代码打印的结果是[3]int,[]int，可以看到前者有长度，后者没有
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	// 通过 make 创建一个切片
	//[]int,指定切片的类型，3是切片的长度，6是切片的容量
	slice3 := make([]int, 3, 6)
	fmt.Println(slice3)

	//给切片赋值
	slice3[0] = 0
	slice3[1] = 1
	slice3[2] = 2

	slice3 = append(slice3, 4)
	// 返回一个新的切片
	fmt.Println(slice3)
}

func TestSlice2(t *testing.T) {

	//我们创建了一个切片，有四个元素，下标命名为0，1，2，3
	slice11 := []int{1, 2, 3, 4}

	// 假设我们要删除下标为0的元素，这段代码的含义是
	// 创建一个空的切片 slice11[:0]
	// 创建一个从下标1到末尾元素的切片 slice11[1:]
	emptySlice := slice11[:0]
	fmt.Println(emptySlice)
	// 给空切片添加元素并返回一个新的切片
	fmt.Println("slice11", slice11)
	slice12 := append(slice11[:0], slice11[1:]...)
	fmt.Println("slice12", slice12)
	fmt.Println("slice11", slice11)

	// 同上我们得出删除下标为i的元素的切片的公式时
	//sliceTemp := append(slice11[:1], slice11[2:]...)
	//fmt.Println(sliceTemp[1])
	//fmt.Println("slice tmp ", sliceTemp)
}

func TestSlice4(t *testing.T) {
	//我们创建了一个切片，有四个元素，下标命名为0，1，2，3
	slice11 := []int{1, 2, 3, 4}
	fmt.Println("TestSlice4 slice11", slice11)
	slice12 := append(slice11[:0], slice11[1:]...)
	fmt.Println("slice12", slice12)
	fmt.Println("slice11", slice11)
}

func TestSlice3(t *testing.T) {

	array11 := [5]int{1, 2, 3, 4, 5}
	slice11 := array11[0:3]
	slice11[0] = 10
	// [10, 2] 和 [4, 5] => [10, 2, 4, 5]
	sliceTemp := append(array11[:2], array11[3:]...)
	fmt.Println("sliceTemp: ", sliceTemp)
	slice11[0] = 11
	fmt.Println(array11, slice11, sliceTemp)
}

func TestSlice6(t *testing.T)  {

	emails := make([]int, 0, 4)
	a := 4
	if a % 2 == 0  {
		emails = append(emails, 10)
	}
	emails = append(emails, 11)
	emails = append(emails, 13)
	fmt.Println(emails[2], len(emails), len(emails) == 2)

	//fmt.Println(a == 4)

	for i, v := range emails {
		fmt.Printf("index: %d, value: %d", i, v)
		fmt.Println()
	}


	var target int
	fmt.Printf("target: %d", target)
	fmt.Println()



}


