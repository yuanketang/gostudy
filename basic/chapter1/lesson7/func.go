package main

import (
	"fmt"
)

func main() {
	fmt.Println("函数")
	// func 函数的名称(函数参数的列表) 函数返回值列表 { // 函数体 }

	// 函数的调用我们使用函数名称(函数的参数) 实参

	_ = add(1, 1)
	fmt.Println("func add = ", add(1, 1))

	// 无参数函数调用方式
	dec()

	// 匿名函数

	bb := func(a int, b int) int {
		return a + b
	}
	fmt.Println("匿名函数调用", bb(2, 2))

	fmt.Println("匿名函数调用方式2", func(a int, b int) int {
		return a + b
	}(3, 3))
	/*
		func(a int, b int) int {
			return a + b
		}(3, 3)
	*/

	opFunc := op("+")
	fmt.Println("匿名函数+", opFunc(1, 9))
	opFunc = op("-")
	fmt.Println("匿名函数-", opFunc(10, 1))

	// 值传递和引用传递

	t1, t2 := 5, 10
	r1, r2 := swap(t1, t2)
	fmt.Println(r1, r2)
	fmt.Println(t1, t2)
	// 闭包
	func() {
		t1, t2 = t2, t1
	}()
	fmt.Println("闭包", t1, t2)

	// 函数的可变参数
	test(1, 2, 3, 4, 5, 6)
}

// 形参
func add(a int, b int) (r int) {
	return a + b
}

func dec() {
	fmt.Println("func dec 被调用")
}

func swap(a int, b int) (int, int) {
	return b, a
}

// 形参
func test(params ...int) {
	fmt.Println(params)
	test2(params...)
}
func test2(params ...int) {

}

// func add2(a, b int, c int64) (r int) {
// 	return a + b
// }

// + - *
func op(s string) func(int, int) int {
	switch {
	case s == "+":
		return func(a, b int) int {
			return a + b
		}
	case s == "-":
		return func(a, b int) int {
			return a - b
		}
	default:
		panic("操作不支持")
	}
}
