package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// CreateIngressRuleParams 安全组入网规则创建参数
type CreateIngressRuleParams struct {
	// TODO
}

// CreateIngressRuleResponse 安全组入网规则创建返回数据
type CreateIngressRuleResponse struct {
	// TODO
}

// CreateIngressRule 安全组入网规则创建
func (s *ECS) CreateIngressRule(p *CreateIngressRuleParams) (resp *CreateIngressRuleResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/sgr/:id/ingress")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
