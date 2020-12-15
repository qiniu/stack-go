package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeEIPsParams 弹性公网IP列表参数
type DescribeEIPsParams struct {
	// TODO
}

// DescribeEIPsResponse 弹性公网IP列表返回数据
type DescribeEIPsResponse struct {
	// TODO
}

// DescribeEIPs 弹性公网IP列表
func (s *VPC) DescribeEIPs(p *DescribeEIPsParams) (resp *DescribeEIPsResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/eip")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
