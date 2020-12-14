package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// CreateVPCParams 专有网络创建参数
type CreateVPCParams struct {
	// TODO
}

// CreateVPCResponse 专有网络创建返回数据
type CreateVPCResponse struct {
	// TODO
}

// CreateVPC 专有网络创建
func (s *VPC) CreateVPC(p *CreateVPCParams) (resp *CreateVPCResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/vpc/:id")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
