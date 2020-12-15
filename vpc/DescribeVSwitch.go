package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeVSwitchParams 虚拟交换机详情参数
type DescribeVSwitchParams struct {
	// TODO
}

// DescribeVSwitchResponse 虚拟交换机详情返回数据
type DescribeVSwitchResponse struct {
	// TODO
}

// DescribeVSwitch 虚拟交换机详情
func (s *VPC) DescribeVSwitch(p *DescribeVSwitchParams) (resp *DescribeVSwitchResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/vswitch/:id")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
