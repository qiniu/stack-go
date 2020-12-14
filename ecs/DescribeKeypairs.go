package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeKeypairsParams 密钥对列表参数
type DescribeKeypairsParams struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	KeypairName        *string `json:"key_pair_name"`
	KeypairFingerPrint *string `json:"key_pair_finger_print"`
	SourceDiskType     *string `json:"source_disk_type"`
}

// DescribeKeypairsResponse 密钥对列表返回数据
type DescribeKeypairsResponse struct {
	Page      int       `json:"page"`
	Size      int       `json:"size"`
	Total     int       `json:"total"`
	RequestID string    `json:"request_id"`
	Data      []Keypair `json:"data"`
}

// DescribeKeypairs 密钥对列表
func (s *ECS) DescribeKeypairs(p *DescribeKeypairsParams) (resp *DescribeKeypairsResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/keypair").WithQueries(p)
	err = s.client.Call(req, &resp)
	return
}
