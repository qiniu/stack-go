package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// AttachKeypairParams 密钥对绑定参数
type AttachKeypairParams struct {
	Name      string `json:"name"` // 密钥对名称
	Instances []struct {
		InstanceID string `json:"instance_id"` // 实例 ID
		RegionID   string `json:"region_id"`   // 区域 ID
	} `json:"instances"`
	Password *string `json:"password,omitempty"` // 资源池 2 主机运行状态下绑定密钥需要输入密码
}

// AttachKeypairResponse 密钥对绑定返回数据
type AttachKeypairResponse struct {
	RequestID string `json:"request_id"`
}

// AttachKeypair 密钥对绑定
func (s *ECS) AttachKeypair(p *AttachKeypairParams) (resp *AttachKeypairResponse, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/v1/vm/keypair/%s/attach", p.Name)).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
