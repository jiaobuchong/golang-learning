package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

// 分块读取文件
//fileread3 -fpath=/path-of-file/test.txt
func main() {
	fptr := flag.String("fpath", "ioread/test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	// 缓冲读取器，buffered reader
	r := bufio.NewReader(f)
	//创建了长度和容量为 3 的字节切片
	b := make([]byte, 3)
	for {
		_, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b))
	}
}
