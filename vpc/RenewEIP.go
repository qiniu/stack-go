package vpc

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/params"
)

// RenewEIPParams 弹性公网IP续费参数
type RenewEIPParams struct {
	RegionID     string  `json:"region_id"`
	AllocationID string  `json:"instance_id"`
	ClientToken  *string `json:"client_token,omitempty"`

	*params.CostParams
}

// RenewEIPResponse 弹性公网IP续费返回数据
type RenewEIPResponse struct {
	RequestID string `json:"request_id"`
	Data      struct {
		OrderHash string  `json:"order"`
		CFee      float64 `json:"c_fee"`
	} `json:"data"`
}

// RenewEIP 弹性公网IP续费
func (s *VPC) RenewEIP(p *RenewEIPParams) (resp *RenewEIPResponse, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/v1/vm/eip/%s/renew", p.AllocationID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
