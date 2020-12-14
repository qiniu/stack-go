package bss

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// CancelOrderParams 取消订单参数
type CancelOrderParams struct {
	OrderHash string `json:"order_hash"`
}

// CancelOrderResponse 取消订单返回数据
type CancelOrderResponse struct {
	RequestID string `json:"request_id"`
}

// CancelOrder 取消订单
func (s *BSS) CancelOrder(p *CancelOrderParams) (resp *CancelOrderResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/order/cancel").WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
