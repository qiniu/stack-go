package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// AttachSecurityGroupParams 安全组绑定参数
type AttachSecurityGroupParams struct {
	// TODO
}

// AttachSecurityGroupResponse 安全组绑定返回数据
type AttachSecurityGroupResponse struct {
	// TODO
}

// AttachSecurityGroup 安全组绑定
func (s *ECS) AttachSecurityGroup(p *AttachSecurityGroupParams) (resp *AttachSecurityGroupResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/sgr/:id/attach")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
