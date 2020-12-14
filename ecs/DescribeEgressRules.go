package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeEgressRulesParams 安全组出网规则列表参数
type DescribeEgressRulesParams struct {
	// TODO
}

// DescribeEgressRulesResponse 安全组出网规则列表返回数据
type DescribeEgressRulesResponse struct {
	// TODO
}

// DescribeEgressRules 安全组出网规则列表
func (s *ECS) DescribeEgressRules(p *DescribeEgressRulesParams) (resp *DescribeEgressRulesResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/sgr/:id/egress")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
