package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// AssociateEIPParams 弹性公网IP绑定参数
type AssociateEIPParams struct {
	// TODO
}

// AssociateEIPResponse 弹性公网IP绑定返回数据
type AssociateEIPResponse struct {
	// TODO
}

// AssociateEIP 弹性公网IP绑定
func (s *VPC) AssociateEIP(p *AssociateEIPParams) (resp *AssociateEIPResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/eip/:id/associate")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
