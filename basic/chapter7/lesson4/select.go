package main

import (
	"fmt"
)

func main() {
	fmt.Println("select")

	ch1 := make(chan int)
	go func() {
		sum := 0
		for i := 1; i < 1000; i++ {
			sum += i
			ch1 <- sum
		}
		close(ch1)
	}()

	ch2 := make(chan int)
	go func() {
		for i := 1; i < 1000; i++ {
			ch2 <- i + 2
		}
		close(ch2)
	}()

	// 1.for
	// for {
	// 	value, ok := <-ch1
	// 	if ok {
	// 		fmt.Println("for ch1 = ", value)
	// 	} else {
	// 		break
	// 	}
	// }

	// 2 for range
	// for c := range ch1 {
	// 	fmt.Println("range ch1 = ", c)
	// }

	// switch类似
	// select 当多个分支条件满足时，会随机选择一个分支执行
	for {
		select {
		// 必须是一个channel的IO操作，要么读，要么写
		case value, ok := <-ch1:
			if ok {
				fmt.Println("select ch1 = ", value)
			}
		case value, ok := <-ch2:
			if ok {
				fmt.Println("select ch2 = ", value)
			}
		}
	}

	// fmt.Println("main结束")
}
