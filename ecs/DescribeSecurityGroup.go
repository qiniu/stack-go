package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeSecurityGroupParams 安全组详情参数
type DescribeSecurityGroupParams struct {
	SecurityGroupID string `json:"security_group_id"`
	RegionID        string `json:"region_id"`
}

// DescribeSecurityGroupResponse 安全组详情返回数据
type DescribeSecurityGroupResponse struct {
	RequestID string `json:"request_id"`
	Data      struct {
		SecurityGroupID   string            `json:"security_group_id"`
		SecurityGroupName string            `json:"security_group_name"`
		RegionID          string            `json:"region_id"`
		Description       string            `json:"description"`
		InnerAccessPolicy InnerAccessPolicy `json:"inner_access_policy"` //enum  Accept | Drop
		Permissions       struct {
			Permission []PermissionType `json:"permission"`
		} `json:"permissions"`
		VPCID string `json:"vpc_id"`
	} `json:"data"`
}

// DescribeSecurityGroup 安全组详情
func (s *ECS) DescribeSecurityGroup(p *DescribeSecurityGroupParams) (resp *DescribeSecurityGroupResponse, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/v1/vm/sgr/%s", p.SecurityGroupID)).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
