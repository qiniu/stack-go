package vpc

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// ModifyEIPSpecParams 弹性公网IP变配参数
type ModifyEIPSpecParams struct {
	RegionID     string `json:"region_id"`     // EIP 所在的地域
	AllocationID string `json:"allocation_id"` // 弹性公网 IP 的 ID
	Bandwidth    string `json:"bandwidth"`     // EIP 的带宽峰值，取值：1~500，单位为 Mbps
}

// ModifyEIPSpecResponse 弹性公网IP变配返回数据
type ModifyEIPSpecResponse struct {
	RequestID string `json:"request_id"` // 请求 ID
}

// ModifyEIPSpec 弹性公网IP变配
//
// 预付费 EIP 调用该接口会生成订单，需要支付后才能变配生效
func (s *VPC) ModifyEIPSpec(p *ModifyEIPSpecParams) (resp *ModifyEIPSpecResponse, err error) {
	req := client.NewRequest(http.MethodPut, fmt.Sprintf("/v1/vm/eip/%s/spec", p.AllocationID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
