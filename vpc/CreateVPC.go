package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// CreateVPCParams 专有网络创建参数
type CreateVPCParams struct {
	RegionID string         `json:"region_id"`
	Vpc      CreateVpc      `json:"vpc"`
	Switch   *CreateVSwitch `json:"switch,omitempty"` // 可选
}

// CreateVPCResponse 专有网络创建返回数据
type CreateVPCResponse struct {
	RequestID string `json:"request_id"`
	Data      struct {
		VPCID        string `json:"vpc_id"`         // 创建的VPC的ID。
		VRouterID    string `json:"v_router_id"`    // 创建VPC后，系统自动创建的路由器的ID。
		RouteTableID string `json:"route_table_id"` // 创建VPC后，系统自动创建的路由表的ID。
	} `json:"data"`
}

// CreateVPC 专有网络创建
func (s *VPC) CreateVPC(p *CreateVPCParams) (resp *CreateVPCResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/vpc").WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
