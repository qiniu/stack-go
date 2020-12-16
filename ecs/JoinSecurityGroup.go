package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// JoinSecurityGroupParams 安全组绑定参数
type JoinSecurityGroupParams struct {
	RegionID        string `json:"region_id"`
	InstanceID      string `json:"instance_id"`
	SecurityGroupID string `Leave:"security_group_id"` // 安全组ID。您可以调用DescribeSecurityGroups查看您可用的安全组。
}

// JoinSecurityGroupResponse 安全组绑定返回数据
type JoinSecurityGroupResponse struct {
	RequestID string `json:"request_id"`
}

// JoinSecurityGroup 安全组绑定
func (s *ECS) JoinSecurityGroup(p *JoinSecurityGroupParams) (resp *JoinSecurityGroupResponse, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/v1/vm/instance/%s/security_group/join", p.InstanceID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
