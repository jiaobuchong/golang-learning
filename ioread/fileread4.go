package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

//逐行读取文件
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
	s := bufio.NewScanner(f)

	//Scan() 方法读取文件的下一行，如果可以读取，就可以使用 Text() 方法
	for s.Scan() {
		fmt.Println(s.Text())
	}
	//当 Scan 返回 false 时，除非已经到达文件末尾（此时 Err() 返回 nil），
	//否则 Err() 就会返回扫描过程中出现的错误。
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
