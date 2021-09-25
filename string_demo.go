package main

import "fmt"

func printBytes(s string) {
	// 这里获取的是字节
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

// rune 是 go 语言的内建类型，它也是 int32 的别称，
// rune 代表一个代码点，代码点无论占用多少个字节，都可以用一个 rune 来表示
func printChars(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func printChars1(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Println()
}

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func main() {
	// go 语言的字符串是一个字节切片
	name := "Hello World"
	printBytes(name)
	printChars1("Señor")
	printChars("Señor")

	// 字符串本身是不可变的
	h := "hello" // h 这个字符串是没有修改的，转换成 rune 之后可以修改了
	fmt.Println(mutate([]rune(h)))
}
