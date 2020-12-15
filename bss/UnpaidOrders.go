package bss

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/params"
)

// UnpaidOrdersParams 未支付订单列表参数
type UnpaidOrdersParams struct {
	ResourceType *params.ResourceType `json:"resource_type"`
}

// UnpaidOrdersResponse 未支付订单列表返回数据
type UnpaidOrdersResponse struct {
	RequestID string `json:"request_id"`
}

// UnpaidOrders 未支付订单列表
func (s *BSS) UnpaidOrders(p *UnpaidOrdersParams) (resp *UnpaidOrdersResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/order/unpaid_orders").WithQueries(p)
	err = s.client.Call(req, &resp)
	return
}
