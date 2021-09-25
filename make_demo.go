package main  // 只有可执行的程序的包名应当为 main，包用于代码的封装和重用，这里的包名是 main
import "fmt"   // 引入 fmt 包

// go 目录是工作区
// src 目录存放源代码文件
func main() {

    // func make（[]T，len，cap）[]T 通过传递类型，长度和容量来创建切片。
    // 容量是可选参数, 默认值为切片长度。make 函数创建一个数组，并返回引用该数组的切片。
    i := make([]int, 5, 5)
    fmt.Println(i)
}