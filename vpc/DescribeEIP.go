package vpc

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeEIPParams 弹性公网IP详情参数
type DescribeEIPParams struct {
	AllocationID string `json:"allocation_id"` // 弹性公网 IP 的 ID
	RegionID     string `json:"region_id"`     // EIP 所在的地域
}

// DescribeEIPResponse 弹性公网IP详情返回数据
type DescribeEIPResponse struct {
	RequestID string  `json:"request_id"` // 请求 ID
	Data      EIPInfo `json:"data"`       // EIP 详情
}

// DescribeEIP 弹性公网IP详情
func (s *VPC) DescribeEIP(p *DescribeEIPParams) (resp *DescribeEIPResponse, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/v1/vm/eip/%s", p.AllocationID)).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
