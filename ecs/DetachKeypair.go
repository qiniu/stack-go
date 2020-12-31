package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DetachKeypairParams 密钥对解绑参数
type DetachKeypairParams struct {
	Name      string `json:"name"` // 密钥对名称
	Instances []struct {
		InstanceID string `json:"instance_id"` // 实例 ID
		RegionID   string `json:"region_id"`   // 区域 ID
	} `json:"instances"` // 实例
}

// DetachKeypairResponse 密钥对解绑返回数据
type DetachKeypairResponse struct {
	RequestID string `json:"request_id"`
}

// DetachKeypair 密钥对解绑
func (s *ECS) DetachKeypair(p *DetachKeypairParams) (resp *DetachKeypairResponse, err error) {
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/vm/keypair/%s/attach", p.Instances)).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
