package main

import (
	"fmt"
	"sync"
)

// 轻量级的线程，Go语言运行时调度的

func Test(n int) {
	defer wg.Done()
	fmt.Println(n)
}

var wg sync.WaitGroup

func main() {
	fmt.Println("并发编程")
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go Test(i)
	}
	wg.Wait()
}
