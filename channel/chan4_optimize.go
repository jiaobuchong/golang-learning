package main

import (
	"fmt"
)

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	// 关闭信道
	close(dchnl)
}

func calcSquares1(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes1(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	// 消费信道里的数据
	for digit := range dch {
		sum += digit * digit * digit
	}
	// 将结果写入信道
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares1(number, sqrch)
	go calcCubes1(number, cubech)

	// 获取信道中的数据
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}
