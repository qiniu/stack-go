package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeKeypairsParams 密钥对列表参数
type DescribeKeypairsParams struct {
	// TODO
}

// DescribeKeypairsResponse 密钥对列表返回数据
type DescribeKeypairsResponse struct {
	// TODO
}

// DescribeKeypairs 密钥对列表
func (s *ECS) DescribeKeypairs(p *DescribeKeypairsParams) (resp *DescribeKeypairsResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/keypair")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
