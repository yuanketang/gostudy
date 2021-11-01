package main

import "fmt"

type MyInt int // 自定义类型

type MyInt2 = int // 类型别名

// byte rune

type Logger struct {
	Desc string
}

func (logger Logger) PrintLog() {
	fmt.Println("我是日志：", logger.Desc)
}

type MyLogger Logger

func (logger MyLogger) PrintPrettyLog() {
	fmt.Printf("我是日志：\t %s \t!!!\n", logger.Desc)
}

func main() {
	fmt.Println("自定义类型与类型别名")

	// 自定义类型
	var a MyInt
	var b MyInt2

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// 旧Logger
	logger := Logger{}
	logger.Desc = "哈哈哈"
	logger.PrintLog()

	fmt.Println("自定义类型拓展后的Logger")
	var myLogger MyLogger
	myLogger.Desc = "哈哈哈"
	myLogger.PrintPrettyLog()

}
