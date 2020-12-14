package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeVPCsParams 专有网络列表参数
type DescribeVPCsParams struct {
	// TODO
}

// DescribeVPCsResponse 专有网络列表返回数据
type DescribeVPCsResponse struct {
	// TODO
}

// DescribeVPCs 专有网络列表
func (s *VPC) DescribeVPCs(p *DescribeVPCsParams) (resp *DescribeVPCsResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/vpc")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
