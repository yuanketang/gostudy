package main

import (
	"fmt"
	"os"
)

func Test() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	panic("error")
	fmt.Println(5)
}

// defer 延迟执行函数是在函数返回前return执行后能够执行的函数
// defer 调用次序是一种前进后出也可以说是后进先出来执行的
// defer 参数是在defer定义时就已经计算好的

func main() {
	fmt.Println("defer延迟调用")

	// Test()
	// for i := 1; i <= 5; i++ {
	// 	defer fmt.Println(i)
	// }

	// defer fmt.Println("hello world")
	// os.Exit(0)
	WriteFile("test2.txt", "Hello world")
	WriteFile("test2.txt", "Hello world2")
}

func WriteFile(fileName string, content string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(content)
	//asdasdasd
}
