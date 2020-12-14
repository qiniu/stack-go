package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// ModifyEgressRuleParams 安全组出网规则修改参数
type ModifyEgressRuleParams struct {
	// TODO
}

// ModifyEgressRuleResponse 安全组出网规则修改返回数据
type ModifyEgressRuleResponse struct {
	// TODO
}

// ModifyEgressRule 安全组出网规则修改
func (s *ECS) ModifyEgressRule(p *ModifyEgressRuleParams) (resp *ModifyEgressRuleResponse, err error) {
	req := client.NewRequest(http.MethodPut, "/v1/vm/sgr/:id/egress")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
