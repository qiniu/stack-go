package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeEIPParams 弹性公网IP创建参数
type DescribeEIPParams struct {
	// TODO
}

// DescribeEIPResponse 弹性公网IP创建返回数据
type DescribeEIPResponse struct {
	// TODO
}

// DescribeEIP 弹性公网IP创建
func (s *VPC) DescribeEIP(p *DescribeEIPParams) (resp *DescribeEIPResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/eip")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
