package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeInstanceParams 主机详情参数
type DescribeInstanceParams struct {
	// TODO
}

// DescribeInstanceResponse 主机详情返回数据
type DescribeInstanceResponse struct {
	// TODO
}

// DescribeInstance 主机详情
func (s *ECS) DescribeInstance(p *DescribeInstanceParams) (resp *DescribeInstanceResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/instance/:id")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
