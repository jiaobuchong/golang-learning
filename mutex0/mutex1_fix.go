package main

import (
	"fmt"
	"sync"
)

var x1 = 0

func increment1(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x1 = x1 + 1
	m.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup

	// 推荐使用 Mutex 来解决竞态条件
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment1(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x1)
}
