package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeVPCsParams 专有网络列表参数
type DescribeVPCsParams struct {
	Page      int     `json:"page"`
	Size      int     `json:"size"`
	RegionID  *string `json:"region_id"`
	VPCID     *string `json:"vpc_id"`
	VPCName   *string `json:"vpc_name"`
	CIDRBlock *string `json:"cidr_block"`
}

// DescribeVPCsResponse 专有网络列表返回数据
type DescribeVPCsResponse struct {
	Page      int       `json:"page"`
	Size      int       `json:"size"`
	Total     int       `json:"total"`
	RequestID string    `json:"request_id"`
	Data      []VPCInfo `json:"data"`
}

// DescribeVPCs 专有网络列表
func (s *VPC) DescribeVPCs(p *DescribeVPCsParams) (resp *DescribeVPCsResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/vpc").WithRegionID(p.RegionID).WithQueries(p)
	err = s.client.Call(req, &resp)
	return
}
