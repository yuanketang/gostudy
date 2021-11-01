package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("计时器和定时器")
	// 计时器
	ticker := time.NewTicker(time.Second * 5)

	// 定时器, 只会执行一次
	timer := time.NewTimer(time.Second * 3)

	timeout := make(chan bool)
	go func() {
		<-time.After(time.Second * 8)
		timeout <- true
	}()
	for {
		select {
		case <-ticker.C:
			fmt.Println("上报服务器数据...")
		case <-timer.C:
			fmt.Println("3秒后执行...")
		case <-timeout:
			fmt.Println("8秒超时...")
		default:
			fmt.Println("服务器运行中...")
			time.Sleep(time.Second * 2)
		}
	}
}
