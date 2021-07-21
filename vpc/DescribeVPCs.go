package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeVPCsParams 专有网络列表参数
type DescribeVPCsParams struct {
	Page     int     `json:"page"`      // 页码。默认：1
	Size     int     `json:"size"`      // 分页大小，最大100。默认20
	RegionID *string `json:"region_id"` // 地域 ID。您可以调用 DescribeRegions 查看最新的七牛云地域列表。
	VPCID    *string `json:"vpc_id"`    // VPC ID
	VPCName  *string `json:"vpc_name"`  // VPC 名称

	// VPC的网段
	//
	// 您可以使用以下网段或其子集作为VPC的网段：
	// - 172.16.0.0/12（默认值）。
	// - 10.0.0.0/8。
	// - 192.168.0.0/16。
	CIDRBlock *string `json:"cidr_block"`
}

// DescribeVPCsResponse 专有网络列表返回数据
type DescribeVPCsResponse struct {
	Page      int    `json:"page"`       // 页码
	Size      int    `json:"size"`       // 分页大小
	Total     int    `json:"total"`      // 安全组总量
	RequestID string `json:"request_id"` // 请求 ID
	Data      []Info `json:"data"`       // VPC 列表
}

// DescribeVPCs 专有网络列表
func (s *VPC) DescribeVPCs(p *DescribeVPCsParams) (resp *DescribeVPCsResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/vpc").WithRegionID(p.RegionID).WithQueries(p)
	err = s.client.Call(req, &resp)
	return
}
