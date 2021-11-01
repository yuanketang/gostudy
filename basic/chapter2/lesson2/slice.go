package main

import "fmt"

func main() {
	fmt.Println("Slice 切片")

	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6} // 数组
	var _ []int                               // 切片， slice

	// 切片只对数组的局部或全部的一种引用形式

	fmt.Println("[1:3]", arr[1:3]) // 半开半闭区间
	fmt.Println("[:3]", arr[:3])
	fmt.Println("[1:]", arr[1:])
	fmt.Println("[:]", arr[:])

	s1 := arr[1:3]

	// fmt.Println("s1[2]", s1[2])
	s2 := s1[2:5]

	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)

	// 切片的长度 len()
	fmt.Println("len(s1) = ", len(s1))
	fmt.Println("len(s2) = ", len(s2))

	// 切片的容量 cap()
	fmt.Println("cap(s1) = ", cap(s1))
	fmt.Println("cap(s2) = ", cap(s2))

	// 切片的操作
	// 1切片的拓展追加 append
	d := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("d", d)
	d1 := d[:3]
	d1 = append(d1, 101)
	fmt.Println("append之后")
	fmt.Println("d1", d1)
	fmt.Println("d", d)

	// d1 = append(d1, 102,103)
	d1 = append(d1, []int{102, 103}...)
	fmt.Println("d1", d1)
	fmt.Println("d", d)

	d1 = append(d1, []int{103, 104}...)
	fmt.Println("d1", d1)

	// 复制 copy()
	f := [...]int{1, 2, 3, 4, 5, 6}
	f1 := f[1:3]
	f2 := f[3:]
	fmt.Println("f1", f1)
	fmt.Println("f2", f2)
	copy(f2, f1)
	fmt.Println("copy之后")
	fmt.Println("f1", f1)
	fmt.Println("f2", f2)
	fmt.Println("f", f)

	// 切片 make()

	t := make([]int, 3, 6)
	t2 := make([]int, 3)
	// t2 默认以2的倍数方式增长

	fmt.Println("t", t, len(t), cap(t))
	fmt.Println("t2", t2, len(t2), cap(t2))
}
