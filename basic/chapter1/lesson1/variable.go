// 定义文件所在包的名字 main 可执行文件
package main

import (
	"fmt"
)

var x = 5

// test := 5

// 单行注释使用//开头
// 多行注释使用/**/

/*
hello
world
ni hao
*/
// main函数是main包下特殊函数，代表可执行应用程序的入口
func main() {
	// 1.单个变量的定义
	// 1.1 基础定义 var 变量名称 变量数据类型
	// 1.2 编译器自动推导类型 var 变量名称 = 初始值
	// 1.3 简洁语法
	var name string = "张三"
	var name1 = "李四"
	var name2 string

	// ****
	name3 := "小红"

	fmt.Printf("%s, %s, %q, %s", name, name1, name2, name3)

	// 多个变量的定义
	// 1.一行定义多个变量
	// 2.多行，分组定义()
	var age, age2 int = 18, 20
	var sex1, sex2 string
	var birth1, birth2 = "2000-01-01", "2000-01-02"

	fmt.Println(age, age2, sex1, sex2, birth1, birth2)
	var (
		a      = 5
		b bool = false
		c string
	)

	fmt.Println(a, b, c)

	fmt.Println(x)
}
