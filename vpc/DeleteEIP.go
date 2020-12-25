package vpc

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteEIPParams 弹性公网IP删除参数
type DeleteEIPParams struct {
	RegionID     string `json:"region_id"`     // EIP 地域 ID
	AllocationID string `json:"allocation_id"` // 弹性公网 IP 的 ID
}

// DeleteEIPResponse 弹性公网IP删除返回数据
type DeleteEIPResponse struct {
	RequestID string `json:"request_id"` // 请求 ID
}

// DeleteEIP 弹性公网IP删除
func (s *VPC) DeleteEIP(p *DeleteEIPParams) (resp *DeleteEIPResponse, err error) {
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/vm/eip/%s", p.AllocationID)).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
