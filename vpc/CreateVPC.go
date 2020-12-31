package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// CreateVPCParams 专有网络创建参数
type CreateVPCParams struct {
	RegionID string         `json:"region_id"`        // 地域 ID
	Vpc      CreateVpc      `json:"vpc"`              // 专有网络信息
	Switch   *CreateVSwitch `json:"switch,omitempty"` // 交换机信息（可选）
}

// CreateVPCResponse 专有网络创建返回数据
type CreateVPCResponse struct {
	RequestID string `json:"request_id"` // 请求 ID
	Data      struct {
		VPCID        string `json:"vpc_id"`         // VPC ID
		VRouterID    string `json:"v_router_id"`    // 路由器 ID
		RouteTableID string `json:"route_table_id"` // 路由表 ID
	} `json:"data"`
}

// CreateVPC 专有网络创建
func (s *VPC) CreateVPC(p *CreateVPCParams) (resp *CreateVPCResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/vpc").WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
