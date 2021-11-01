package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("常量的定义")

	// 常量的定义，我们使用const
	// 单个常量的定义
	const pi = 3.14
	const os_version string = "1.0.0"

	fmt.Println(pi, os_version)
	// 单行定义多个常量
	const a, b = 3, 4
	fmt.Println(a, b)
	// var d, e = 3, 4

	const (
		c = 5
		d = 10
	)
	fmt.Println(c, d)

	// 计算三角函数

	cc := math.Sqrt(a*a + b*b)
	fmt.Println("三角函数", cc)

	// 枚举 enum
	// 其他语言习惯常量或枚举我们使用大写 final String OS_VERSION = "123"
	// 首字符大写默认public其他包可见
	const (
		langC    = "C"
		langJava = "JAVA"
		langGo   = "GO"
	)
	fmt.Println(langC, langJava, langGo)

	// 自增类型的枚举
	// 每一次const iota初始化为0

	const (
		zz = iota
		zz1
	)
	const (
		tt = iota
		tt1
		tt2
		_
		tt4
	)

	fmt.Println(zz, zz1, tt, tt1, tt2, tt4)

	// 简洁语法_占位符
	name1, _ := "张三", "李四"
	fmt.Println(name1)
}
