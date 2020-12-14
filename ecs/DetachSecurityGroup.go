package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DetachSecurityGroupParams 安全组解绑参数
type DetachSecurityGroupParams struct {
	// TODO
}

// DetachSecurityGroupResponse 安全组解绑返回数据
type DetachSecurityGroupResponse struct {
	// TODO
}

// DetachSecurityGroup 安全组解绑
func (s *ECS) DetachSecurityGroup(p *DetachSecurityGroupParams) (resp *DetachSecurityGroupResponse, err error) {
	req := client.NewRequest(http.MethodDelete, "/v1/vm/sgr/:id/attach")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
