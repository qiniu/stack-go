package vpc

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeVSwitchParams 虚拟交换机详情参数
type DescribeVSwitchParams struct {
	VSwitchID string `json:"vswitch_id"`
	RegionID  string `json:"region_id"`
}

// DescribeVSwitchResponse 虚拟交换机详情返回数据
type DescribeVSwitchResponse struct {
	RequestID string      `json:"request_id"`
	Data      VSwitchInfo `json:"data"`
}

// DescribeVSwitch 虚拟交换机详情
func (s *VPC) DescribeVSwitch(p *DescribeVSwitchParams) (resp *DescribeVSwitchResponse, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/v1/vm/switch/%s", p.VSwitchID)).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
