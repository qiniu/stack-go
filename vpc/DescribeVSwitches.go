package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeVSwitchesParams 虚拟交换机列表参数
type DescribeVSwitchesParams struct {
	// TODO
}

// DescribeVSwitchesResponse 虚拟交换机列表返回数据
type DescribeVSwitchesResponse struct {
	// TODO
}

// DescribeVSwitches 虚拟交换机列表
func (s *VPC) DescribeVSwitches(p *DescribeVSwitchesParams) (resp *DescribeVSwitchesResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/vswitch")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
