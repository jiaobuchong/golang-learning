package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

//使用值接受者声明的方法，既可以用值来调用，也能用指针调用

func (p Person) Describe() {  // 使用值接受者实现
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() {  // 使用指针接受者实现
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func main() {
	var d1 Describer
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Describe()
	p2 := Person{"James", 32}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"Washington", "USA"}

	/* 如果下面一行取消注释会导致编译错误：
	   cannot use a (type Address) as type Describer
	   in assignment: Address does not implement
	   Describer (Describe method has pointer
	   receiver)
	*/

	// a 属于值类型，并没有实现 Describer 接口
	// 对于使用指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的。参考 method4.go
	// a 这个值类型并没有实现接口，还有一种解释：a 赋值给 d2，d2 中存储的具体值不能取到地址，也就没法调用
	// 	a.Describe() 是可以调用的，因为可以自动解地址 (&a).Describe()
	//d2 = a
	// d2 是指针类型，底层应该知道如何解地址
	a.Describe()

	d2 = &a // 这是合法的
	// 因为在第 22 行，Address 类型的指针实现了 Describer 接口
	d2.Describe()

}
