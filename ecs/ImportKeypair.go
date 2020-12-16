package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// ImportKeypairParams 密钥对导入参数
type ImportKeypairParams struct {
	KeypairName   string `json:"key_pair_name"`
	PublicKeyBody string `json:"public_key_body"` // 为空时代表创建，否则是导入
}

// ImportKeypairResponse 密钥对导入返回数据
type ImportKeypairResponse struct {
	RequestID string `json:"request_id"`
}

// ImportKeypair 密钥对导入
func (s *ECS) ImportKeypair(p *ImportKeypairParams) (resp *ImportKeypairResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/keypair").WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
