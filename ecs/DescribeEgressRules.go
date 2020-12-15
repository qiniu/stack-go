package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeEgressRulesParams 安全组出网规则列表参数
type DescribeEgressRulesParams struct {
	SecurityGroupID string `json:"security_group_id"`
	RegionID        string `json:"region_id"`
	Page            int    `json:"page"`
	Size            int    `json:"size"`
}

// DescribeEgressRulesResponse 安全组出网规则列表返回数据
type DescribeEgressRulesResponse struct {
	Page      int              `json:"page"`
	Size      int              `json:"size"`
	Total     int              `json:"total"`
	RequestID string           `json:"request_id"`
	Data      []PermissionType `json:"data"`
}

// DescribeEgressRules 安全组出网规则列表
func (s *ECS) DescribeEgressRules(p *DescribeEgressRulesParams) (resp *DescribeEgressRulesResponse, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/v1/vm/sgr/%s/egress", p.SecurityGroupID)).WithQueries(p).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
