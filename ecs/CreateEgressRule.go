package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// CreateEgressRuleParams 安全组出网规则创建参数
type CreateEgressRuleParams struct {
	// TODO
}

// CreateEgressRuleResponse 安全组出网规则创建返回数据
type CreateEgressRuleResponse struct {
	// TODO
}

// CreateEgressRule 安全组出网规则创建
func (s *ECS) CreateEgressRule(p *CreateEgressRuleParams) (resp *CreateEgressRuleResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/sgr/:id/egress")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
