package vpc

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// ModifyEIPSpecParams 弹性公网IP变配参数
type ModifyEIPSpecParams struct {
	RegionID     string `json:"region_id"`
	AllocationID string `json:"allocation_id"`
	Bandwidth    string `json:"bandwidth"`
}

// ModifyEIPSpecResponse 弹性公网IP变配返回数据
type ModifyEIPSpecResponse struct {
	RequestID string `json:"request_id"`
}

// ModifyEIPSpec 弹性公网IP变配
// 预付费EIP调用该接口会生成订单，需要支付后才能变配生效
func (s *VPC) ModifyEIPSpec(p *ModifyEIPSpecParams) (resp *ModifyEIPSpecResponse, err error) {
	req := client.NewRequest(http.MethodPut, fmt.Sprintf("/v1/vm/eip/%s/spec", p.AllocationID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
