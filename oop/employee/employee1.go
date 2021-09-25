package employee

import "fmt"

// 定义成 employee 之后，要在其它包里就不可以访问了
// 结构体不可引用
// 隐藏所有不可引用的结构体的所有字段
// 也是 go 的最佳实践，不会在外部包需要 employee 的字段，因此我们也让这些字段无法引用
type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}


//该函数接收必要的参数，返回一个新创建的 employee 结构体变量
func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	e := employee {firstName, lastName, totalLeave, leavesTaken}
	return e
}


func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}

func main() {

}
