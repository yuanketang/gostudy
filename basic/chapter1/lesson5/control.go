package main

import (
	"fmt"
)

func main() {
	fmt.Println("条件语句")
	// 需求计算订单优惠价
	// 1.订单金额小于等于100打九折
	// 2.订单金额大于100小于等于300打八折
	// 3.大于300打五折
	// if else if else
	total := 150.00 // 浮点数字面量
	discount := 0.0
	if total <= 100 {
		// 1.订单金额小于等于100打九折
		discount = total * 0.9
	} else if total > 100 && total <= 300 {
		// 2.订单金额大于100小于等于300打八折
		discount = total * 0.8
	} else {
		// 3.大于300打五折
		discount = total * 0.5
	}

	fmt.Printf("if 订单优惠后的价格：%.2f\n", discount)

	// switch

	switch {
	case total <= 100:
		discount = total * 0.9
	case total > 100 && total <= 300:
		discount = total * 0.8
	default:
		discount = total * 0.5
	}

	fmt.Printf("switch 订单优惠后的价格：%.2f", discount)
}
