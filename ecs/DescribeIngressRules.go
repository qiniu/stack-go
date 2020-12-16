package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeIngressRulesParams 安全组入网规则列表参数
type DescribeIngressRulesParams struct {
	SecurityGroupID string `json:"security_group_id"`
	RegionID        string `json:"region_id"`
	Page            int    `json:"page"`
	Size            int    `json:"size"`
}

// DescribeIngressRulesResponse 安全组入网规则列表返回数据
type DescribeIngressRulesResponse struct {
	Page      int              `json:"page"`
	Size      int              `json:"size"`
	Total     int              `json:"total"`
	RequestID string           `json:"request_id"`
	Data      []PermissionType `json:"data"`
}

// DescribeIngressRules 安全组入网规则列表
func (s *ECS) DescribeIngressRules(p *DescribeIngressRulesParams) (resp *DescribeIngressRulesResponse, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/v1/vm/sgr/%s/ingress", p.SecurityGroupID)).WithQueries(p).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
