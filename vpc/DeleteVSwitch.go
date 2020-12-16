package vpc

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteVSwitchParams 虚拟交换机删除参数
type DeleteVSwitchParams struct {
	VSwitchID string `json:"vswitch_id"`
	RegionID  string `json:"region_id"`
}

// DeleteVSwitchResponse 虚拟交换机删除返回数据
type DeleteVSwitchResponse struct {
	RequestID string `json:"request_id"`
}

// DeleteVSwitch 虚拟交换机删除
func (s *VPC) DeleteVSwitch(p *DeleteVSwitchParams) (resp *DeleteVSwitchResponse, err error) {
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/vm/switch/%s", p.VSwitchID)).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
