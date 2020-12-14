package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteVPCParams 专有网络删除参数
type DeleteVPCParams struct {
	// TODO
}

// DeleteVPCResponse 专有网络删除返回数据
type DeleteVPCResponse struct {
	// TODO
}

// DeleteVPC 专有网络删除
func (s *VPC) DeleteVPC(p *DeleteVPCParams) (resp *DeleteVPCResponse, err error) {
	req := client.NewRequest(http.MethodDelete, "/v1/vm/vpc/:id")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
