package main

import (
	"fmt"
	"time"
)

func main() {
	// 秒
	fmt.Println(time.Now().Unix())

	// 毫秒

	fmt.Println("millis ", time.Now())
	time.Sleep(1)

	cur := time.Now().Sub(time.Unix(1631478143, 0))
	fmt.Println(cur > 24*time.Hour)

	date := time.Now().Format("2006")
	fmt.Println(date)



	 //
}
