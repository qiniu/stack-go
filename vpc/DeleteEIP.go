package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteEIPParams 弹性公网IP删除参数
type DeleteEIPParams struct {
	// TODO
}

// DeleteEIPResponse 弹性公网IP删除返回数据
type DeleteEIPResponse struct {
	// TODO
}

// DeleteEIP 弹性公网IP删除
func (s *VPC) DeleteEIP(p *DeleteEIPParams) (resp *DeleteEIPResponse, err error) {
	req := client.NewRequest(http.MethodDelete, "/v1/vm/eip/:id")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
