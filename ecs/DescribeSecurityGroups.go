package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeSecurityGroupsParams 安全组列表参数
type DescribeSecurityGroupsParams struct {
	Page              int     `json:"page"`                // 页码。默认：1
	Size              int     `json:"size"`                // 分页大小，最大100。默认20
	RegionID          *string `json:"region_id"`           // 地域 ID。您可以调用 DescribeRegions 查看最新的七牛云地域列表。
	SecurityGroupID   *string `json:"security_group_id"`   // 安全组 ID
	SecurityGroupName *string `json:"security_group_name"` // 安全组名称
	VPCID             *string `json:"vpc_id"`              // VPC ID
}

// DescribeSecurityGroupsResponse 安全组列表返回数据
type DescribeSecurityGroupsResponse struct {
	Page      int             `json:"page"`       // 页码
	Size      int             `json:"size"`       // 分页大小
	Total     int             `json:"total"`      // 安全组总量
	RequestID string          `json:"request_id"` // 请求 ID
	Data      []SecurityGroup `json:"data"`       // 安全组列表
}

// DescribeSecurityGroups 安全组列表
func (s *ECS) DescribeSecurityGroups(p *DescribeSecurityGroupsParams) (resp *DescribeSecurityGroupsResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/sgr").WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
