package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeVSwitchesParams 虚拟交换机列表参数
type DescribeVSwitchesParams struct {
	Page     int     `json:"page"`      // 页码。默认：1
	Size     int     `json:"size"`      // 分页大小，最大100。默认20
	RegionID *string `json:"region_id"` // 地域 ID。您可以调用 DescribeRegions 查看最新的七牛云地域列表
	VPCID    *string `json:"vpc_id"`    // VPC ID
	ZoneID   *string `json:"zone_id"`   // 可用区 ID。您可以通过调用DescribeZones接口获取可用区ID
}

// DescribeVSwitchesResponse 虚拟交换机列表返回数据
type DescribeVSwitchesResponse struct {
	Page      int           `json:"page"`       // 页码
	Size      int           `json:"size"`       // 分页大小
	Total     int           `json:"total"`      // 虚拟交换机总量
	RequestID string        `json:"request_id"` // 请求 ID
	Data      []VSwitchInfo `json:"data"`       // 虚拟交换机列表
}

// DescribeVSwitches 虚拟交换机列表
func (s *VPC) DescribeVSwitches(p *DescribeVSwitchesParams) (resp *DescribeVSwitchesResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/switch").WithRegionID(p.RegionID).WithQueries(p)
	err = s.client.Call(req, &resp)
	return
}
