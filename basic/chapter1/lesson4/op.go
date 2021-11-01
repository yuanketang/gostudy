package main

import (
	"fmt"
)

func main() {
	fmt.Println("基本运算和运算符")

	// 1 + 1 = 2
	// 知识点
	// 1.运算符
	// 2.表达式
	// 3.双目运算符
	// 4.字面量 3.5
	// 5.表达式是可以求值的
	// (1 + 0) + (2 -1) = 5-3
	// x / x + y / y = 2
	// var name string = "Hello World"

	fmt.Println("1 + 1 = ", 1+1)

	// 拓展变量的定义与常数的定义

	var a int = 5
	var b int = a * 4
	const c = 4
	var d int = c * a * b
	const e = c * 5
	fmt.Printf("a = %d, b = %d, d = %d, e = %d\n", a, b, d, e)

	// 算术运算符+ - * / %

	// 整数的除法结果还是整数
	fmt.Println(5/2, 5.0/2)

	fmt.Println(a % c)

	// fmt.Println(5 + 4i % 4)
	// 关系运算符== != > < >= <=

	// 逻辑运算符&& || !
	// && 逻辑与操作
	var age int = 20
	fmt.Println(age >= 20 && age < 100)
	// || 逻辑或
	fmt.Println(age >= 20 || age < 15)
	// ！非运算符(单目运算符) 1 + 1
	fmt.Println(!(age > 20))
	// 位运算符& | ^ << >>
	// & 按位与 | 按位或 ^按位异或 << 左移运算符 >> 右移运算符
	fmt.Println(1 << 2)
	kb := 1 << (10 * 1)
	fmt.Println(kb)
	fmt.Println(4 >> 2)
	// 赋值运算符= += -= /= %= <<= >>= &= ^= |=
	var x int = 10
	// x = x + 1
	x += 1
	fmt.Println(x)
}
