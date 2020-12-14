package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// UnAssociateEIPParams 弹性公网IP解绑参数
type UnAssociateEIPParams struct {
	// TODO
}

// UnAssociateEIPResponse 弹性公网IP解绑返回数据
type UnAssociateEIPResponse struct {
	// TODO
}

// UnAssociateEIP 弹性公网IP解绑
func (s *VPC) UnAssociateEIP(p *UnAssociateEIPParams) (resp *UnAssociateEIPResponse, err error) {
	req := client.NewRequest(http.MethodDelete, "/v1/vm/eip/:id/associate")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
