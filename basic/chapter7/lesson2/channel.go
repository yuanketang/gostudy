package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("channel")
	// 1.无缓冲的通道，同步通道
	// var ch chan int
	// fmt.Println("未初始化", ch)
	// ch = make(chan int)
	// fmt.Println("初始化完成", ch)
	ch := make(chan int)
	go func() {
		// 从通道接收数据
		fmt.Println("从通道中接收数据", <-ch)
	}()
	// 向通道发送数据
	ch <- 100

	// 2 有缓冲的通道，不是同步通道
	ch2 := make(chan int, 3)
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	// ch2 <- 4 // 缓冲区满的时候不能发送数据
	fmt.Println("从ch2中接收数据", <-ch2)
	fmt.Println("从ch2中接收数据", <-ch2)
	fmt.Println("从ch2中接收数据", <-ch2)

	// fmt.Println("从ch2中接收数据", <-ch2) 缓冲区为空时不能读取数据
	close(ch2)
	// ch2 <- 1 // 不能向已关闭的通道发送数据
	fmt.Println("从ch2中接收数据", <-ch2) // 从关闭通道中缓冲区为0的通道返回数据会返回通道元素类型的默认0值

	// close(ch2) // 不能再次关闭已经关闭的通道

	// 模拟比特币挖矿

	blockChain := make(chan string)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go Miner(i, blockChain)
	}
	wg.Add(1)
	go func(ch chan string) {
		defer wg.Done()
		for c := range ch {
			fmt.Println(c)
		}
	}(blockChain)
	wg.Wait()
	fmt.Println("挖矿成功完成")
}

const MagicNum = "66"

func Miner(n int, ch chan string) {
	defer wg.Done()
	num := rand.Int()
	numStr := strconv.Itoa(num)
	if strings.HasPrefix(numStr, MagicNum) {
		ch <- fmt.Sprintf("矿工编号 %d, 挖矿成功 %d", n, num)
	}
	// } else {
	// 	fmt.Printf("矿工编号 %d, 计算hash %d\n", n, num)
	// }
}
