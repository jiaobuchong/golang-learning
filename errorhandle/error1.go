package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt")

	// 类型断言
	//类型断言格式： y, ok := x.(T)
	//err, ok := err.(*os.PathError) //成功：f为os.PathError，ok为true，否则ok为false
	//ok值通常立刻用于决定是否执行下一步，惯用法：
	//if f, ok := w.(*os.File); ok { // ... use f ... }

	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}
