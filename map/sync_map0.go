package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对

	//使用 Range 配合一个回调函数进行遍历操作，
	//通过回调函数返回内部遍历出来的值，
	//Range 参数中回调函数的返回值在需要继续迭代遍历时，
	//返回 true，终止迭代遍历时，返回 false。
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
