package main

import (
	"fmt"
)

// int string bool 基本数据类型

type user struct {
	id   int
	name string
	addr address
}

type address struct {
	mobile string
	city   string
}

// 匿名成员的数据类型相同的只允许出现一次
type teacher struct {
	int
	string
	city string
}

func main() {
	fmt.Println("结构体")

	// 方式1
	var u user
	fmt.Printf("%#v\n", u)

	// 方式2
	u1 := user{id: 1, name: "小张"}
	fmt.Printf("u1 %T, %#v\n", u1, u1)

	// 方式3
	u2 := &user{2, "小李", address{mobile: "123", city: "北京"}}
	fmt.Printf("u2 %T, %#v\n", u2, u2)

	// 方式4 new()
	u3 := new(user) // &user{}
	u3.id = 888
	fmt.Printf("u3 %T, %#v\n", u3, u3)

	// 匿名结构体
	var u4 struct {
		id   int
		name string
	}
	fmt.Printf("u4 %T, %#v\n", u4, u4)

	// 结构体如何访问成员
	u4.id = 101
	u4.name = "刘德华"
	fmt.Printf("u4 %T, %#v\n", u4, u4)

	// 结构体的嵌套

	u5 := user{}
	u5.id = 666
	u5.name = "小孙"
	u5.addr.mobile = "188"
	u5.addr.city = "广州"

	fmt.Printf("u5 %T, %#v\n", u5, u5)

	// 很少用
	t := teacher{}
	t.int = 999
	t.string = "孙老师"
	t.city = "上海"
}
