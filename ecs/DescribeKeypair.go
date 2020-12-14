package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeKeypairParams 密钥对详情参数
type DescribeKeypairParams struct {
	Name string `json:"name"`
}

// DescribeKeypairResponse 密钥对详情返回数据
type DescribeKeypairResponse struct {
	RequestID string  `json:"request_id"`
	Data      Keypair `json:"data"`
}

// DescribeKeypair 密钥对详情
func (s *ECS) DescribeKeypair(p *DescribeKeypairParams) (resp *DescribeKeypairResponse, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/v1/vm/keypair/%s", p.Name))
	err = s.client.Call(req, &resp)
	return
}
