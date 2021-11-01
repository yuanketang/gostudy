package main

import (
	"fmt"
	"gostudy/chapter3/lesson2/bird"
)

func main() {
	fmt.Println("结构体")

	b := bird.Bird1{}
	b.Name = "花花"
	b.Fly2()
}
