package main

import "fmt"

// interface definition
type VowelsFinder interface {
	// FindVowels 方法
	FindVowels() []rune
}

// 创建一个 MyString 类型
type MyString string

// MyString implements VowelsFinder
// 类型包含接口中的方法，也就实现了该接口
func (ms MyString) FindVowels() []rune {
	//如果一个类型包含了接口中声明的所有方法，那么它就隐式地实现了 Go 接口。
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

// 接口是方法签名的集合（method signature）
// 当一个类型定义了接口中的所有方法，称它实现了该接口
// 接口指定了一个类型应该具有的方法，并由该类型决定如何实现这些方法。
func main() {
	name := MyString("Sam Anderson")
	var v VowelsFinder
	// MyString 实现了 VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())

}
