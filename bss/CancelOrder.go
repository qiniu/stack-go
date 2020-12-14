package bss

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// CancelOrderParams 取消订单参数
type CancelOrderParams struct {
	// TODO
}

// CancelOrderResponse 取消订单返回数据
type CancelOrderResponse struct {
	// TODO
}

// CancelOrder 取消订单
func (s *BSS) CancelOrder(p *CancelOrderParams) (resp *CancelOrderResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/order/cancel")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
