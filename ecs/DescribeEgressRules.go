package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeEgressRulesParams 安全组出网规则列表参数
type DescribeEgressRulesParams struct {
	SecurityGroupID string `json:"security_group_id"` // 安全组ID。您可以调用 DescribeSecurityGroups 查看您可用的安全组。
	RegionID        string `json:"region_id"`         // 安全组所属地域ID。您可以调用 DescribeRegions 查看最新的七牛云地域列表。
	Page            int    `json:"page"`              // 页码。默认：1
	Size            int    `json:"size"`              // 分页大小，最大100。默认20
}

// DescribeEgressRulesResponse 安全组出网规则列表返回数据
type DescribeEgressRulesResponse struct {
	Page      int              `json:"page"`       // 页码
	Size      int              `json:"size"`       // 分页大小
	Total     int              `json:"total"`      // 安全组总量
	RequestID string           `json:"request_id"` // 请求 ID
	Data      []PermissionType `json:"data"`       // 安全组权限规则
}

// DescribeEgressRules 安全组出网规则列表
func (s *ECS) DescribeEgressRules(p *DescribeEgressRulesParams) (resp *DescribeEgressRulesResponse, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/v1/vm/sgr/%s/egress", p.SecurityGroupID)).WithQueries(p).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
