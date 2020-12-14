package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeInstanceParams 主机详情参数
type DescribeInstanceParams struct {
	RegionID   string `json:"region_id"`
	InstanceID string `json:"instance_id"`
}

// DescribeInstanceResponse 主机详情返回数据
type DescribeInstanceResponse struct {
	RequestID string   `json:"request_id"`
	Data      Instance `json:"data"`
}

// DescribeInstance 主机详情
func (s *ECS) DescribeInstance(p *DescribeInstanceParams) (resp *DescribeInstanceResponse, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/v1/vm/instance/%s", p.InstanceID)).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
