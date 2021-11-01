package main

import "fmt"

func main() {
	fmt.Println("指针")
	// 指针，是用来保存内存地址的一种特殊的变量类型

	var p *int
	var p1 *float32
	var c int

	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(c)

	fmt.Printf("p 类型 %T\n", p)
	fmt.Printf("p %p\n", p)
	// nil 指针，函数，管道，切片，interface, map

	// & 取地址
	name := 10
	p = &name
	fmt.Printf("p %p\n", p)

	fmt.Println(name)

	// * 解指针操作符
	fmt.Println(*p)

	*p = 100
	fmt.Println(name)

	var p3 *int
	p3 = new(int)
	*p3 = 101
	fmt.Println(*p3)

}
