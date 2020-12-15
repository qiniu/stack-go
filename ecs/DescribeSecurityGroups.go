package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeSecurityGroupsParams 安全组列表参数
type DescribeSecurityGroupsParams struct {
	Page              int     `json:"page"`
	Size              int     `json:"size"`
	RegionID          *string `json:"region_id"`
	SecurityGroupID   *string `json:"security_group_id"`
	SecurityGroupName *string `json:"security_group_name"`
	VPCID             *string `json:"vpc_id"`
}

// DescribeSecurityGroupsResponse 安全组列表返回数据
type DescribeSecurityGroupsResponse struct {
	Page      int             `json:"page"`
	Size      int             `json:"size"`
	Total     int             `json:"total"`
	RequestID string          `json:"request_id"`
	Data      []SecurityGroup `json:"data"`
}

// DescribeSecurityGroups 安全组列表
func (s *ECS) DescribeSecurityGroups(p *DescribeSecurityGroupsParams) (resp *DescribeSecurityGroupsResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/sgr").WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
