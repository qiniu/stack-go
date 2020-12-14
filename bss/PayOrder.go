package bss

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// PayOrderParams 支付订单参数
type PayOrderParams struct {
	// TODO
}

// PayOrderResponse 支付订单返回数据
type PayOrderResponse struct {
	// TODO
}

// PayOrder 支付订单
func (s *BSS) PayOrder(p *PayOrderParams) (resp *PayOrderResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/order/pay")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
