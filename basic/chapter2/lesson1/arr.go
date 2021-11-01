package main

import "fmt"

func main() {
	fmt.Println("数组")

	var arr [3]int

	fmt.Println(arr)

	var arr1 [3]int = [3]int{}

	fmt.Println(arr1)

	// 简洁语法定义

	arr2 := [5]int{1, 3, 5, 7, 9}

	fmt.Println(arr2)

	arr3 := [...]int{2, 4, 6}

	fmt.Println(arr3)

	// len

	fmt.Println("arr3 length", len(arr3))

	// 多维数组的定义

	t2 := [3][5]int{{1}, {3}, {5, 4, 3, 2, 1}}

	fmt.Println("t2 length", t2)

	// for
	// arr2 := [5]int{1, 3, 5, 7, 9}
	for i := 0; i < len(arr2); i++ {
		fmt.Println("i = ", i, arr2[i])
	}

	// for range
	fmt.Println("for range")
	for _, v := range arr2 {
		fmt.Println("i = ", v)
	}
	for i := range arr2 {
		fmt.Println("i = ", i)
	}

	// 数组是值类型
	var test [3]int = [3]int{1, 3, 5}
	// var test2 [5]int = [5]int{1, 3, 5, 6, 7}
	printArray(&test)
	fmt.Println("循环后", test)

}

func printArray(arr *[3]int) {
	fmt.Println("循环开始")
	arr[0] = 101
	for _, v := range arr {
		fmt.Println("i = ", v)
	}
}
