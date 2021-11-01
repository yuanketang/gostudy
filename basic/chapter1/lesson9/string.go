package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("字符与字符串")

	s := "你好，hello"
	// len()
	fmt.Println("s 长度", len(s))

	// 字符串底层使用字节来保存
	// len返回字符串字节数

	// byte 字符
	var c byte = 'h'
	fmt.Printf("c = %c\n", c)
	// rune unicode的字符
	var ch rune = '中'
	fmt.Printf("ch = %c\n", ch)

	fmt.Println("*s 长度", utf8.RuneCountInString(s))

	for _, r := range s {
		fmt.Printf("rune %T, %U, %c\n", r, r, r)
	}
}
