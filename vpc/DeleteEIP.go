package vpc

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteEIPParams 弹性公网IP删除参数
type DeleteEIPParams struct {
	RegionID     string `json:"region_id"`
	AllocationID string `json:"allocation_id"`
}

// DeleteEIPResponse 弹性公网IP删除返回数据
type DeleteEIPResponse struct {
	RequestID string `json:"request_id"`
}

// DeleteEIP 弹性公网IP删除
func (s *VPC) DeleteEIP(p *DeleteEIPParams) (resp *DeleteEIPResponse, err error) {
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/vm/eip/%s", p.AllocationID)).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
