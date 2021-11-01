package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a bool      = false
		b string    = "张三"
		c int       = -16
		d uint      = 100
		e float32   = 3.14
		f complex64 = 3 + 2i
	)
	fmt.Println(a, b, c, d, e, f)

	// 强制类型转换
	var var1 int32 = 100
	var2 := int64(var1)

	fmt.Println("var2 = ", var2)

	// sqrt(a*a + b*b) = c

	var aa, bb = 3, 4
	cc := math.Sqrt(float64(aa*aa + bb*bb))
	fmt.Println("cc = ", cc)

	// IEEE754
}
