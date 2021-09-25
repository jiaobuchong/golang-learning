package main

import (
	"fmt"
	"os"
)

// 将字节写入文件
func main() {
	f, err := os.Create("/Users/bj89109ml/go/src/hello/iowrite/one_file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	n2, err := f.Write(d2)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(n2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
