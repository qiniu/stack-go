package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeVPCParams 专有网络详情参数
type DescribeVPCParams struct {
	// TODO
}

// DescribeVPCResponse 专有网络详情返回数据
type DescribeVPCResponse struct {
	// TODO
}

// DescribeVPC 专有网络详情
func (s *VPC) DescribeVPC(p *DescribeVPCParams) (resp *DescribeVPCResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/vpc/:id")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
