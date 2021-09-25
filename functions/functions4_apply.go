package main

// 头等函数的实际用途
import (
	"fmt"
)

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

//函数接收一个 students 切片和一个函数作为参数，这个函数会计算一个学生是否满足筛选条件。
//filter 的第二个参数是一个函数。这个函数接收 student 参数，返回一个 bool 值。
func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}

	// 根据成绩过滤
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)

	// 根据地址过滤
	c := filter(s, func(s student) bool {
		if s.country == "India" {
			return true
		}
		return false
	})
	fmt.Println(c)
}
