package main

// 面向对象——多态
// go 通过接口来实现多态。我们已经讨论过，在 Go 语言中，我们是隐式地实现接口。
// 一个类型如果定义了接口所声明的全部方法，那它就实现了该接口。

//所有实现了接口的 类型，都可以把它的值保存在一个 接口类型的变量 中。
//在 Go 中，我们使用接口的这种特性来实现多态。

import (
	"fmt"
)

type Income interface {
	// 计算返回项目的收入
	calculate() int
	// 返回项目名称
	source() string
}

// 项目的结构体类型
type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

// 新增一个收益流
type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}

//有 FixedBilling 和 TimeAndMaterial 两种类型
// ic []Income 表现了多态
// 新增一个收益流（income stream）不需要改 calculateNetIncome 的代码，这个就是多态带来的好处
func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d", netincome)
}

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)
}
