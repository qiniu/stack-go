package bss

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// UnpaidOrdersParams 未支付订单列表参数
type UnpaidOrdersParams struct {
	// TODO
}

// UnpaidOrdersResponse 未支付订单列表返回数据
type UnpaidOrdersResponse struct {
	// TODO
}

// UnpaidOrders 未支付订单列表
func (s *BSS) UnpaidOrders(p *UnpaidOrdersParams) (resp *UnpaidOrdersResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/order/unpaid_orders")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
