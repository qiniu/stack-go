package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/params"
)

// RenewInstanceParams 主机续费参数
type RenewInstanceParams struct {
	*params.CostParams
	RegionID   string `json:"region_id"`
	InstanceID string `json:"instance_id"`

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。
	ClientToken *string `json:"client_token"`
}

// RenewInstanceResponse 主机续费返回数据
type RenewInstanceResponse struct {
	RequestID string `json:"request_id"`
	Data      struct {
		OrderHash string  `json:"order"`
		CFee      float64 `json:"c_fee"`
	} `json:"data"`
}

// RenewInstance 主机续费
func (s *ECS) RenewInstance(p *RenewInstanceParams) (resp *RenewInstanceResponse, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/v1/vm/instance/%s/renew", p.InstanceID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
