package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeVSwitchesParams 虚拟交换机列表参数
type DescribeVSwitchesParams struct {
	Page     int     `json:"page"`
	Size     int     `json:"size"`
	RegionID *string `json:"region_id"`
	VPCID    *string `json:"vpc_id"`
	ZoneID   *string `json:"zone_id"`
}

// DescribeVSwitchesResponse 虚拟交换机列表返回数据
type DescribeVSwitchesResponse struct {
	Page      int           `json:"page"`
	Size      int           `json:"size"`
	Total     int           `json:"total"`
	RequestID string        `json:"request_id"`
	Data      []VSwitchInfo `json:"data"`
}

// DescribeVSwitches 虚拟交换机列表
func (s *VPC) DescribeVSwitches(p *DescribeVSwitchesParams) (resp *DescribeVSwitchesResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/switch").WithRegionID(p.RegionID).WithQueries(p)
	err = s.client.Call(req, &resp)
	return
}
