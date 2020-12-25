package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeEIPsParams 弹性公网IP列表参数
type DescribeEIPsParams struct {
	Page           int        `json:"page"`            // 页码。默认：1
	Size           int        `json:"size"`            // 分页大小，最大100。默认20
	RegionID       *string    `json:"region_id"`       // 地域 ID。您可以调用 DescribeRegions 查看最新的七牛云地域列表
	AllocationID   *string    `json:"allocation_id"`   // 要查询的 EIP 实例的 ID
	AllocationName *string    `json:"allocation_name"` // 要查询的 EIP 实例的名称
	IPAddress      *string    `json:"ip_address"`      // 要查询的 EIP 实例的 IP
	Status         *EIPStatus `json:"status"`          // EIP 状态: Creating, Associating, Unassociating, InUse, Available
	InstanceID     *string    `json:"instance_id"`     // 实例 ID
}

// DescribeEIPsResponse 弹性公网IP列表返回数据
type DescribeEIPsResponse struct {
	Page      int       `json:"page"`       // 页码
	Size      int       `json:"size"`       // 分页大小
	Total     int       `json:"total"`      // EIP 总量
	RequestID string    `json:"request_id"` // 请求 ID
	Data      []EIPInfo `json:"data"`       // EIP 列表
}

// DescribeEIPs 弹性公网IP列表
func (s *VPC) DescribeEIPs(p *DescribeEIPsParams) (resp *DescribeEIPsResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/eip").WithRegionID(p.RegionID).WithQueries(p)
	err = s.client.Call(req, &resp)
	return
}
