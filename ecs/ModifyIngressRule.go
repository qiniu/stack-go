package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// ModifyIngressRuleParams 安全组入网规则修改参数
type ModifyIngressRuleParams struct {
	// TODO
}

// ModifyIngressRuleResponse 安全组入网规则修改返回数据
type ModifyIngressRuleResponse struct {
	// TODO
}

// ModifyIngressRule 安全组入网规则修改
func (s *ECS) ModifyIngressRule(p *ModifyIngressRuleParams) (resp *ModifyIngressRuleResponse, err error) {
	req := client.NewRequest(http.MethodPut, "/v1/vm/sgr/:id/ingress")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
