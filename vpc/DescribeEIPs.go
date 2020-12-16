package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeEIPsParams 弹性公网IP列表参数
type DescribeEIPsParams struct {
	Page           int        `json:"page"`
	Size           int        `json:"size"`
	RegionID       *string    `json:"region_id"`
	AllocationID   *string    `json:"allocation_id"`
	AllocationName *string    `json:"allocation_name"`
	IPAddress      *string    `json:"ip_address"`
	Status         *EIPStatus `json:"status"`
	InstanceID     *string    `json:"instance_id"`
}

// DescribeEIPsResponse 弹性公网IP列表返回数据
type DescribeEIPsResponse struct {
	Page      int       `json:"page"`
	Size      int       `json:"size"`
	Total     int       `json:"total"`
	RequestID string    `json:"request_id"`
	Data      []EIPInfo `json:"data"`
}

// DescribeEIPs 弹性公网IP列表
func (s *VPC) DescribeEIPs(p *DescribeEIPsParams) (resp *DescribeEIPsResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/eip").WithRegionID(p.RegionID).WithQueries(p)
	err = s.client.Call(req, &resp)
	return
}
