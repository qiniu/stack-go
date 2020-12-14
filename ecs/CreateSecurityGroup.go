package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// CreateSecurityGroupParams 安全组创建参数
type CreateSecurityGroupParams struct {
	// TODO
}

// CreateSecurityGroupResponse 安全组创建返回数据
type CreateSecurityGroupResponse struct {
	// TODO
}

// CreateSecurityGroup 安全组创建
func (s *ECS) CreateSecurityGroup(p *CreateSecurityGroupParams) (resp *CreateSecurityGroupResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/sgr")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
