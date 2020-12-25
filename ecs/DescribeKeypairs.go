package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeKeypairsParams 密钥对列表参数
type DescribeKeypairsParams struct {
	Page               *int    `json:"page"`                  // 页码。默认：1
	Size               *int    `json:"size"`                  // 分页大小，最大100。默认20
	KeypairName        *string `json:"key_pair_name"`         // 密钥对名称
	KeypairFingerPrint *string `json:"key_pair_finger_print"` // 密钥对的指纹
}

// DescribeKeypairsResponse 密钥对列表返回数据
type DescribeKeypairsResponse struct {
	Page      int       `json:"page"`       // 页码
	Size      int       `json:"size"`       // 分页大小
	Total     int       `json:"total"`      // 密钥对总量
	RequestID string    `json:"request_id"` // 请求 ID
	Data      []Keypair `json:"data"`       // 密钥对列表
}

// DescribeKeypairs 密钥对列表
func (s *ECS) DescribeKeypairs(p *DescribeKeypairsParams) (resp *DescribeKeypairsResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/keypair").WithQueries(p)
	err = s.client.Call(req, &resp)
	return
}
