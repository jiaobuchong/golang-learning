package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("io/test.txt")
	// 使用绝对路径获取文件
	//data, err := ioutil.ReadFile("/home/naveen/go/src/filehandling/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	// 将 data 转换为 string
	fmt.Println("Contents of file:", string(data))
}