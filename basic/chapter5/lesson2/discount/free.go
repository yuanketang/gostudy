package discount

import "fmt"

type Order struct {
	Id    int
	Total float64
}

// 免费会员
type FreeUser struct {
	Order *Order
}

func (u FreeUser) ComputeOrder() {
	fmt.Printf("订单编号：%d, 订单金额：%.2f\n", u.Order.Id, u.Order.Total*0.9)
}
