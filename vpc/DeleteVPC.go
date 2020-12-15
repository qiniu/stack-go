package vpc

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteVPCParams 专有网络删除参数
type DeleteVPCParams struct {
	RegionID string `json:"region_id"`
	VPCID    string `json:"vpc_id"`
}

// DeleteVPCResponse 专有网络删除返回数据
type DeleteVPCResponse struct {
	RequestID string `json:"request_id"`
}

// DeleteVPC 专有网络删除
func (s *VPC) DeleteVPC(p *DeleteVPCParams) (resp *DeleteVPCResponse, err error) {
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/vm/vpc/%s", p.VPCID)).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
