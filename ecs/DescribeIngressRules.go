package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeIngressRulesParams 安全组入网规则列表参数
type DescribeIngressRulesParams struct {
	// TODO
}

// DescribeIngressRulesResponse 安全组入网规则列表返回数据
type DescribeIngressRulesResponse struct {
	// TODO
}

// DescribeIngressRules 安全组入网规则列表
func (s *ECS) DescribeIngressRules(p *DescribeIngressRulesParams) (resp *DescribeIngressRulesResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/sgr/:id/ingress")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
