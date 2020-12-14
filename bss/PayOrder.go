package bss

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// PayOrderParams 支付订单参数
type PayOrderParams struct {
	OrderHash string `json:"order_hash"`
}

// PayOrderResponse 支付订单返回数据
type PayOrderResponse struct {
	RequestID string `json:"request_id"`
}

// PayOrder 支付订单
func (s *BSS) PayOrder(p *PayOrderParams) (resp *PayOrderResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/order/pay").WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
