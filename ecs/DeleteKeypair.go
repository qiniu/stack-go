package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteKeypairParams 密钥对删除参数
type DeleteKeypairParams struct {
	KeypairNames []string `json:"key_pair_names"`
}

// DeleteKeypairResponse 密钥对删除返回数据
type DeleteKeypairResponse struct {
	RequestID string `json:"request_id"`
}

// DeleteKeypair 密钥对删除
func (s *ECS) DeleteKeypair(p *DeleteKeypairParams) (resp *DeleteKeypairResponse, err error) {
	req := client.NewRequest(http.MethodDelete, "/v1/vm/keypair").WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
