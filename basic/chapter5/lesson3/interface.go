package main

import "fmt"

type Printer interface {
	PrintOrder() // 方法的签名或原型
}

type Qr interface {
	Printer
	PrintQr()
}

type Order struct {
	Id    int
	Total float64
}

func (o Order) PrintOrder() {
	fmt.Println(o.Id)
}

type OrderQr struct {
	Order
}

func (o OrderQr) PrintQr() {
	fmt.Println("打印二维码", o.Order.Id)
}

type Empty interface{}

func main() {
	fmt.Println("接口")

	var p Printer = Order{Id: 999, Total: 100.0}
	p.PrintOrder()

	var q Qr = OrderQr{Order{Id: 666, Total: 350.0}}
	q.PrintOrder()
	q.PrintQr()

	// 空接口
	var empty Empty = 456
	empty = false
	fmt.Println(empty)

	var s bool = empty.(bool)
	fmt.Println(s)

	// 空接口应用场景
	map1 := map[string]interface{}{
		"id":   1,
		"name": "小李",
		"sex":  true,
	}

	fmt.Printf("%#v", map1)
}
