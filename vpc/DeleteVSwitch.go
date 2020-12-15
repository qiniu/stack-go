package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteVSwitchParams 虚拟交换机删除参数
type DeleteVSwitchParams struct {
	// TODO
}

// DeleteVSwitchResponse 虚拟交换机删除返回数据
type DeleteVSwitchResponse struct {
	// TODO
}

// DeleteVSwitch 虚拟交换机删除
func (s *VPC) DeleteVSwitch(p *DeleteVSwitchParams) (resp *DeleteVSwitchResponse, err error) {
	req := client.NewRequest(http.MethodDelete, "/v1/vm/vswitch/:id")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
