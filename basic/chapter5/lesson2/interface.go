package main

import (
	"fmt"
	"gostudy/chapter5/lesson2/discount"
)

type CommonUser struct {
	Order *discount.Order
}

func (c CommonUser) ComputeOrder() {
	fmt.Println(c.Order.Id)
}

func (c *CommonUser) PrintCode() {
	fmt.Println("打印二维码", c.Order.Id)
}

func main() {
	fmt.Println("接口")

	order := &discount.Order{
		Id:    999,
		Total: 300.0,
	}
	PrintOrderDiscount(discount.FreeUser{Order: order})
	PrintOrderDiscount(discount.VipUser{Order: order})

	// 接口组合
	fmt.Println("接口组合")
	var dc DiscounterAndCoder
	dc = &CommonUser{
		Order: order,
	}
	dc.ComputeOrder()
	dc.PrintCode()
}

type Discounter interface {
	ComputeOrder()
}

type Coder interface {
	PrintCode()
}

type DiscounterAndCoder interface {
	Discounter
	Coder
}

func PrintOrderDiscount(d Discounter) {
	if _, ok := d.(discount.FreeUser); ok {
		fmt.Println("您享受的是免费会员折扣")
	} else {
		fmt.Println("您享受的是VIP会员折扣")
	}
	d.ComputeOrder()
}
