package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeKeypairParams 密钥对详情参数
type DescribeKeypairParams struct {
	// TODO
}

// DescribeKeypairResponse 密钥对详情返回数据
type DescribeKeypairResponse struct {
	// TODO
}

// DescribeKeypair 密钥对详情
func (s *ECS) DescribeKeypair(p *DescribeKeypairParams) (resp *DescribeKeypairResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/keypair/:id")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
