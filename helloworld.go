package main  // 只有可执行的程序的包名应当为 main，包用于代码的封装和重用，这里的包名是 main
import "fmt"   // 引入 fmt 包


//go 基础教程：https://studygolang.com/subject/2
// https://golangbot.com/learn-golang-series/
func variant1() {
    // 常量
    var age int = 29

    const g = 55   // 常量

    // 简短声明，左边至少有一个尚未声明的变量
    a := true
    b := false

    fmt.Println("hello 世界")
    fmt.Println("age is", age)

    fmt.Println(a, b)
}



// go 目录是工作区
// src 目录存放源代码文件
func main() {
    variant1()
}