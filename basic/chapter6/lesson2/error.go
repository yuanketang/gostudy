package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("错误处理")
	defer func() {
		err := recover()
		if myErr, ok := err.(MyError); ok {
			fmt.Println("main函数捕获错误", myErr.Code, myErr.Reason)
		} else {
			fmt.Println(err)
		}
	}()
	WriteFile("file.txt", "你好，Go")
}

type MyError struct {
	Code   uint
	Reason string
}

func (e MyError) Error() string {
	return fmt.Sprintf("自定义错误码：%d 错误原因：%s", e.Code, e.Reason)
}

//
func WriteFile(fileName string, content string) {
	file, err := os.Open(fileName)
	defer func() {
		err := recover()
		switch err.(type) {
		case *os.PathError:
			fmt.Println("WriteFile捕获错误：文件不存在")
			panic(err)
		default:
			panic(MyError{Code: 1001, Reason: "文件不存在"})
		}
	}()
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(content)
}
