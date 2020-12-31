package vpc

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/params"
)

// RenewEIPParams 弹性公网IP续费参数
type RenewEIPParams struct {
	RegionID     string  `json:"region_id"`              // 地域 ID。您可以调用 DescribeRegions 查看最新的七牛云地域列表
	AllocationID string  `json:"allocation_id"`          // 弹性公网 IP 的 ID
	ClientToken  *string `json:"client_token,omitempty"` // 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。

	*params.CostParams
}

// RenewEIPResponse 弹性公网IP续费返回数据
type RenewEIPResponse struct {
	RequestID string `json:"request_id"` // 请求 ID
	Data      struct {
		OrderHash string  `json:"order"` // 订单号
		CFee      float64 `json:"c_fee"` // 支付金额
	} `json:"data"`
}

// RenewEIP 弹性公网IP续费
func (s *VPC) RenewEIP(p *RenewEIPParams) (resp *RenewEIPResponse, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/v1/vm/eip/%s/renew", p.AllocationID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
