package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// RenewEIPParams 弹性公网IP续费参数
type RenewEIPParams struct {
	// TODO
}

// RenewEIPResponse 弹性公网IP续费返回数据
type RenewEIPResponse struct {
	// TODO
}

// RenewEIP 弹性公网IP续费
func (s *VPC) RenewEIP(p *RenewEIPParams) (resp *RenewEIPResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/eip/:id/renew")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
