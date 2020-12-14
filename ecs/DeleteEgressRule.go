package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteEgressRuleParams 安全组出网规则删除参数
type DeleteEgressRuleParams struct {
	// TODO
}

// DeleteEgressRuleResponse 安全组出网规则删除返回数据
type DeleteEgressRuleResponse struct {
	// TODO
}

// DeleteEgressRule 安全组出网规则删除
func (s *ECS) DeleteEgressRule(p *DeleteEgressRuleParams) (resp *DeleteEgressRuleResponse, err error) {
	req := client.NewRequest(http.MethodDelete, "/v1/vm/sgr/:id/egress")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
