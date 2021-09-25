package main

import (
	"fmt"
	"sync"
)

var (
	queueEmailConfs sync.Map
)

func main() {

	// 从已记录的map中读取
	queue, ok := queueEmailConfs.Load("1")
	// 记录当前游标
	//index := 0
	fmt.Println("ok ", ok)
	if ok {
		fmt.Println("queue ", queue)
	}
	//if ok {
		// 如果能读取到队列相应的配置则获取其游标
		//index = queue.(*emails).Index
	//}

}
