package main

import (
	"context"
	"fmt"
	"time"
)

type MyKey string

type MyValue string

func main() {
	fmt.Println("context")

	// 1 WithCancel
	// 2 WithDeadline
	// 3 WithTimeout
	// 4 WithValue
	// ctx, cancel := context.WithCancel(context.Background())
	// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	// go TestDeadline(ctx)
	// go TestCancel(ctx)
	// 1. runtime.Goexit()
	// 2. channel
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	// go TestTimeout(ctx)

	var myKey MyKey = "key"
	var myValue MyValue = "我是值"
	ctx := context.WithValue(context.Background(), myKey, myValue)
	go TestValue(ctx, myKey)
	<-time.After(time.Second * 10)
	// cancel()
	// <-time.After(time.Second * 2)
	fmt.Println("main结束")
}

// 1
func TestCancel(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("主协程取消了,我是子协程也同时结束了...")
			return
		default:
			fmt.Println("TestCancel 我是子协程，我在运行中..")
			time.Sleep(time.Second)
		}
	}
}

func TestDeadline(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("我是子协程，我运行了5秒就结束了...")
			return
		default:
			fmt.Println("TestDeadline 我是子协程，我在运行中..")
			time.Sleep(time.Second)
		}
	}
}

func TestTimeout(ctx context.Context) {
	<-time.After(time.Second * 5)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("我是子协程，我超时3秒结束了...")
			return
		default:
			fmt.Println("TestTimeout 我是子协程，我在运行中..")
			time.Sleep(time.Second)
		}
	}
}

func TestValue(ctx context.Context, myKey MyKey) {
	for {
		select {
		default:
			fmt.Println("TestValue 我是子协程，我获取到了参数", ctx.Value(myKey))
			time.Sleep(time.Second)
		}
	}
}
