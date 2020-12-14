package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeSecurityGroupParams 安全组详情参数
type DescribeSecurityGroupParams struct {
	// TODO
}

// DescribeSecurityGroupResponse 安全组详情返回数据
type DescribeSecurityGroupResponse struct {
	// TODO
}

// DescribeSecurityGroup 安全组详情
func (s *ECS) DescribeSecurityGroup(p *DescribeSecurityGroupParams) (resp *DescribeSecurityGroupResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/sgr/:id")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
