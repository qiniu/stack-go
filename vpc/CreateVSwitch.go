package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// CreateVSwitchParams 虚拟交换机创建参数
type CreateVSwitchParams struct {
	// TODO
}

// CreateVSwitchResponse 虚拟交换机创建返回数据
type CreateVSwitchResponse struct {
	// TODO
}

// CreateVSwitch 虚拟交换机创建
func (s *VPC) CreateVSwitch(p *CreateVSwitchParams) (resp *CreateVSwitchResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/vswitch/:id")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
