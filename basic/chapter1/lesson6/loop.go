package main

import "fmt"

func main() {
	fmt.Println("循环")
	// ++ 自增操作符，作用是将变量的值自增+1
	// -- 自减操作符，作用是将变量的值自减-1
	a := 5
	a++
	fmt.Println("a++", a)
	a--
	fmt.Println("a--", a)

	// 循环1
	// 循环的初值
	// 循环的结束条件(每次循环之前)
	// 循环增量或减量（循环之后）
	// continue 作用，从continue之后的语句都会被跳过，而执行下次循环
	for i := 1; i <= 10; i++ {
		// 循环体
		if i%2 == 0 {
			break
		}
		fmt.Println("循环1, i = ", i)
	}

	// 循环2
	// for i := 1; i <= 10; {
	// 	// 循环体
	// 	fmt.Println("循环2, i = ", i)
	// 	i++
	// }

	// 无限循环，死循环
	// for {
	// 	// 无限循环，死循环
	// }

}
