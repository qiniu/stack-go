package vpc

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeVPCParams 专有网络详情参数
type DescribeVPCParams struct {
	VPCID    string `json:"vpc_id"`
	RegionID string `json:"region_id"`
}

// DescribeVPCResponse 专有网络详情返回数据
type DescribeVPCResponse struct {
	RequestID string  `json:"request_id"`
	Data      VPCInfo `json:"data"`
}

// DescribeVPC 专有网络详情
func (s *VPC) DescribeVPC(p *DescribeVPCParams) (resp *DescribeVPCResponse, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/v1/vm/vpc/%s", p.VPCID)).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
