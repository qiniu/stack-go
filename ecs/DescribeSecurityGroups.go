package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeSecurityGroupsParams 安全组列表参数
type DescribeSecurityGroupsParams struct {
	// TODO
}

// DescribeSecurityGroupsResponse 安全组列表返回数据
type DescribeSecurityGroupsResponse struct {
	// TODO
}

// DescribeSecurityGroups 安全组列表
func (s *ECS) DescribeSecurityGroups(p *DescribeSecurityGroupsParams) (resp *DescribeSecurityGroupsResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/sgr")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
