package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteIngressRuleParams 安全组入网规则删除参数
type DeleteIngressRuleParams struct {
	// TODO
}

// DeleteIngressRuleResponse 安全组入网规则删除返回数据
type DeleteIngressRuleResponse struct {
	// TODO
}

// DeleteIngressRule 安全组入网规则删除
func (s *ECS) DeleteIngressRule(p *DeleteIngressRuleParams) (resp *DeleteIngressRuleResponse, err error) {
	req := client.NewRequest(http.MethodDelete, "/v1/vm/sgr/:id/ingress")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
