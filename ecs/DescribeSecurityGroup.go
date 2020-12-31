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
	RequestID string `json:"request_id"` // 请求 ID

	Data struct {
		SecurityGroupID   string `json:"security_group_id"`   // 安全组 ID
		SecurityGroupName string `json:"security_group_name"` // 安全组名称
		RegionID          string `json:"region_id"`           // 地域 ID
		Description       string `json:"description"`         // 安全组描述信息

		// 安全组内的ECS实例之间的内网连通策略。
		//
		// 取值范围（取值不区分大小写）：
		// - `Accept`：互通
		// - `Drop`：隔离
		InnerAccessPolicy InnerAccessPolicy `json:"inner_access_policy"`

		// 安全组权限规则
		Permissions struct {
			Permission []PermissionType `json:"permission"`
		} `json:"permissions"`

		VPCID string `json:"vpc_id"` // VPC ID
	} `json:"data"`
}

// DescribeSecurityGroup 安全组详情
func (s *ECS) DescribeSecurityGroup(p *DescribeSecurityGroupParams) (resp *DescribeSecurityGroupResponse, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/v1/vm/sgr/%s", p.SecurityGroupID)).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
