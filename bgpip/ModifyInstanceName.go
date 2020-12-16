package bgpip

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// ModifyInstanceNameParams BGP高防实例名称修改参数
type ModifyInstanceNameParams struct {
	RegionID   string `json:"region_id"`   // 地域 id
	ResourceID string `json:"resource_id"` // 高防实例 id
	Name       string `json:"name"`        // 新名称
}

// ModifyInstanceNameResponse BGP高防实例名称修改返回数据
type ModifyInstanceNameResponse struct {
	RequestID string `json:"request_id"`
}

// ModifyInstanceName BGP高防实例名称修改
func (s *BGPIP) ModifyInstanceName(p *ModifyInstanceNameParams) (resp *ModifyInstanceNameResponse, err error) {
	req := client.NewRequest(http.MethodPatch, fmt.Sprintf("/v1/security/bgpip/resource/%s/name", p.ResourceID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
