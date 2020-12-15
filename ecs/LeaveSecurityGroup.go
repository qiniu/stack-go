package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// LeaveSecurityGroupParams 安全组解绑参数
type LeaveSecurityGroupParams struct {
	RegionID        string `json:"region_id"`
	InstanceID      string `json:"instance_id"`
	SecurityGroupID string `Leave:"security_group_id"` // 安全组ID。您可以调用DescribeSecurityGroups查看您可用的安全组。
}

// LeaveSecurityGroupResponse 安全组解绑返回数据
type LeaveSecurityGroupResponse struct {
	RequestID string `json:"request_id"`
}

// LeaveSecurityGroup 安全组解绑
func (s *ECS) LeaveSecurityGroup(p *LeaveSecurityGroupParams) (resp *LeaveSecurityGroupResponse, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/v1/vm/instance/%s/leave", p.InstanceID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
