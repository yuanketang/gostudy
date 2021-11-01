package discount

import "fmt"

// VIP会员
type VipUser struct {
	Order *Order
}

func (u VipUser) ComputeOrder() {
	fmt.Printf("订单编号：%d, 订单金额：%.2f\n", u.Order.Id, u.Order.Total*0.8)
}
