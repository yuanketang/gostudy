package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("数据竞争与锁")

	var x int32

	// 自增+1
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&x, 1)
		}()
	}

	wg.Wait()
	fmt.Println("计算后 x = ", x)

}
